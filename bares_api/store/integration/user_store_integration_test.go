package integration

import (
	"bares_api/store"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestUserStoreIntegration(t *testing.T) {
	// Inicia o banco de dados
	dbStore, err := StartDatabase()
	if err != nil {
		t.Errorf("Falha ao criar banco de dados e tabelas: %s", err)
	}
	defer dbStore.DatabaseClose() // Garante que a conexão será fechada no final

	// storeUsers
	storeUsers := store.NewUser(dbStore.DB)

	// Cria usuários
	users, err := CreateUsers(storeUsers)
	if err != nil {
		t.Errorf("Falha ao criar banco de dados e tabelas: %s", err)
	}

	// Recuperar e verificar usuários
	for _, user := range users {
		// Testar GetUsuarioByEmail
		retrievedUser, err := storeUsers.GetUserByEmail(user.Email)
		if err != nil {
			t.Errorf("Erro ao recuperar usuário pelo email: %s", err)
		}
		if retrievedUser.Email != user.Email {
			t.Errorf("Usuário recuperado não corresponde ao usuário criado")
		}

		// Testar GetUsuario pela id
		retrievedUser, err = storeUsers.GetUser(user.UsuarioID)
		if err != nil {
			t.Errorf("Erro ao recuperar usuário pelo ID: %s", err)
		}
		if retrievedUser.UsuarioID != user.UsuarioID {
			t.Errorf("Usuário recuperado não corresponde ao usuário criado")
		}

		// Testar UpdateUsuario
		retrievedUser.Nome = "Novo Nome"
		err = storeUsers.UpdateUser(retrievedUser)
		if err != nil {
			t.Errorf("Erro ao atualizar usuário: %s", err)
		}

		// Verificar se o usuário foi atualizado
		updatedUser, err := storeUsers.GetUser(user.UsuarioID)
		if err != nil {
			t.Errorf("Erro ao recuperar usuário pelo ID após atualização: %s", err)
		}
		if updatedUser.Nome != "Novo Nome" {
			t.Errorf("Usuário não foi atualizado corretamente")
		}

		// Testar DeleteUsuario
		err = storeUsers.DeleteUser(user.UsuarioID)
		if err != nil {
			t.Errorf("Erro ao deletar usuário: %s", err)
		}

		// Verificar se o usuário foi deletado
		_, err = storeUsers.GetUser(user.UsuarioID)
		if err == nil {
			t.Errorf("Usuário deveria ter sido deletado")
		}
	}

	// Limpeza final
	if _, err := dbStore.DB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDBName)); err != nil {
		log.Fatal("Falha ao limpar o banco de dados de teste após o teste:", err)
	}
}
