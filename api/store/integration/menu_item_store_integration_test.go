package integration

import (
	"bares_api/store"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestItemMenuStoreIntegration(t *testing.T) {
	// Inicia o banco de dados
	dbStore, err := StartDatabase()
	if err != nil {
		t.Errorf("Falha ao criar banco de dados e tabelas: %s", err)
	}
	defer dbStore.DB.Close()

	// storeItensMenu
	storeItensMenu := store.NewMenuItem(dbStore.DB)

	// Cria itens para o menu
	itens, err := CreateItensMenu(storeItensMenu)
	if err != nil {
		t.Errorf("Falha ao criar banco de dados e tabelas: %s", err)
	}

	// Recuperar e verificar itemMenus
	for _, item := range itens {
		// Testar GetItemMenuByNome
		retrievedItem, err := storeItensMenu.GetMenuItemByName(item.Name)
		if err != nil {
			t.Errorf("Erro ao recuperar itemMenu pelo email: %s", err)
		}
		if retrievedItem.Name != item.Name {
			t.Errorf("ItemMenu recuperado não corresponde ao itemMenu criado")
		}

		// Testar GetItemMenu
		retrievedItem, err = storeItensMenu.GetMenuItem(item.Id)
		if err != nil {
			t.Errorf("Erro ao recuperar itemMenu pelo ID: %s", err)
		}
		if retrievedItem.Id != item.Id {
			t.Errorf("ItemMenu recuperado não corresponde ao itemMenu criado")
		}

		// Testar UpdateItemMenu
		retrievedItem.Name = "Novo Nome"
		err = storeItensMenu.UpdateMenuItem(retrievedItem)
		if err != nil {
			t.Errorf("Erro ao atualizar itemMenu: %s", err)
		}

		// Verificar se o itemMenu foi atualizado
		updatedUser, err := storeItensMenu.GetMenuItem(item.Id)
		if err != nil {
			t.Errorf("Erro ao recuperar itemMenu pelo ID após atualização: %s", err)
		}
		if updatedUser.Name != "Novo Nome" {
			t.Errorf("itemMenu não foi atualizado corretamente")
		}

		// Testar DeleteItemMenu
		err = storeItensMenu.DeleteMenuItem(item.Id)
		if err != nil {
			t.Errorf("Erro ao deletar itemMenu: %s", err)
		}

		// Verificar se o itemMenu foi deletado
		_, err = storeItensMenu.GetMenuItem(item.Id)
		if err == nil {
			t.Errorf("itemMenu deveria ter sido deletado")
		}
	}

	// Limpeza final
	if _, err := dbStore.DB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDBName)); err != nil {
		log.Fatal("Falha ao limpar o banco de dados de teste após o teste:", err)
	}
}
