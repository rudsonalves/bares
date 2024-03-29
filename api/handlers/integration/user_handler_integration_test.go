package integration

import (
	"bares_api/handlers"
	"bares_api/models"
	"bares_api/services"
	"bares_api/store"
	storeIntegration "bares_api/store/integration"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCreateUsuarioIntegration(t *testing.T) {
	log.Print("Iniciando o banco de dados...")
	dbStore, err := storeIntegration.StartDatabase()
	if err != nil {
		t.Errorf("Falha ao iniciar o banco de dados de teste: %v", err)
	}
	defer dbStore.DB.Close()

	// Configurar o service para usar o banco de dados de teste
	log.Print("Configurar o service para usar o banco de dados de teste")
	usuarioStore := store.NewUser(dbStore.DB)
	usuarioSrevice := services.NewUsuarioService(usuarioStore)
	usuarioHandler := handlers.NewUserHandler(usuarioSrevice)

	// Cria um servidor HTTP de teste
	log.Print("Cria um servidor HTTP de teste")
	server := httptest.NewServer(http.HandlerFunc(usuarioHandler.CreateUser))
	defer server.Close()

	// Criar Usuários
	users := UsersGenerate()

	for _, user := range users {
		log.Print("usuario:", user)

		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("Erro ao marshalizar o payload: %v", err)
		}
		log.Print("Payload criado:", string(payload))

		// Cria a requisição de teste
		log.Print("Cria a requisição de teste")
		req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(payload))
		if err != nil {
			t.Fatalf("Erro ao criar a requisição: %v", err)
		}

		// Envia a requisição
		log.Print("Envia a requisição")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Erro ao enviar a requisição: %v", err)
		}
		defer resp.Body.Close()

		// Verifica a resposta
		log.Print("Verifica a resposta")
		if resp.StatusCode != http.StatusCreated {
			bodyBytes, _ := io.ReadAll(resp.Body)
			t.Errorf("Status esperado: %v, recebido: %v, corpo: %s",
				http.StatusCreated,
				resp.StatusCode,
				string(bodyBytes),
			)
		}

		// Decodifica a resposta para verificar se o usuário foi criado corretamente
		log.Print("Decodifica a resposta para verificar se o usuário foi criado corretamente")
		var createdUser models.User
		if err := json.NewDecoder(resp.Body).Decode(&createdUser); err != nil {
			t.Fatalf("Erro ao decodificar a resposta: %v", err)
		}

		// Verificar se o usuário criado corresponde ao enviado
		if createdUser.Email != user.Email {
			t.Errorf("Email esperado: %v, recebido: %v", user.Email, createdUser.Email)
		}
	}

	// Limpeza final
	if _, err := dbStore.DB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbStore.DBName)); err != nil {
		log.Fatal("Falha ao limpar o banco de dados de teste após o teste:", err)
	}
}
