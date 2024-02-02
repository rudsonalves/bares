package integration

import (
	"bares_api/store"
	"fmt"
	"log"
	"testing"
)

func TestItemMenuStoreIntegration(t *testing.T) {
	// Inicia o banco de dados
	dbStore, err := StartDatabase()
	if err != nil {
		t.Errorf("Falha ao criar banco de dados e tabelas: %s", err)
	}
	defer dbStore.DB.Close()

	// storeItensMenu
	storeItensMenu := store.NewItensMenu(dbStore.DB)

	// Cria itens para o menu
	itens, err := CreateItensMenu(storeItensMenu)
	if err != nil {
		t.Errorf("Falha ao criar banco de dados e tabelas: %s", err)
	}

	// Recuperar e verificar itemMenus
	for _, item := range itens {
		// Testar GetItemMenuByNome
		retrievedItem, err := storeItensMenu.GetItemMenuByNome(item.Nome)
		if err != nil {
			t.Errorf("Erro ao recuperar itemMenu pelo email: %s", err)
		}
		if retrievedItem.Nome != item.Nome {
			t.Errorf("ItemMenu recuperado não corresponde ao itemMenu criado")
		}

		// Testar GetItemMenu
		retrievedItem, err = storeItensMenu.GetItemMenu(item.ItemID)
		if err != nil {
			t.Errorf("Erro ao recuperar itemMenu pelo ID: %s", err)
		}
		if retrievedItem.ItemID != item.ItemID {
			t.Errorf("ItemMenu recuperado não corresponde ao itemMenu criado")
		}

		// Testar UpdateItemMenu
		retrievedItem.Nome = "Novo Nome"
		err = storeItensMenu.UpdateItemMenu(retrievedItem)
		if err != nil {
			t.Errorf("Erro ao atualizar itemMenu: %s", err)
		}

		// Verificar se o itemMenu foi atualizado
		updatedUser, err := storeItensMenu.GetItemMenu(item.ItemID)
		if err != nil {
			t.Errorf("Erro ao recuperar itemMenu pelo ID após atualização: %s", err)
		}
		if updatedUser.Nome != "Novo Nome" {
			t.Errorf("itemMenu não foi atualizado corretamente")
		}

		// Testar DeleteItemMenu
		err = storeItensMenu.DeleteItemMenu(item.ItemID)
		if err != nil {
			t.Errorf("Erro ao deletar itemMenu: %s", err)
		}

		// Verificar se o itemMenu foi deletado
		_, err = storeItensMenu.GetItemMenu(item.ItemID)
		if err == nil {
			t.Errorf("itemMenu deveria ter sido deletado")
		}
	}

	// Limpeza final
	if _, err := dbStore.DB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDBName)); err != nil {
		log.Fatal("Falha ao limpar o banco de dados de teste após o teste:", err)
	}
}
