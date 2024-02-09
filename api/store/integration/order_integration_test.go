package integration

import (
	"bares_api/models"
	"bares_api/store"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOrderIntegration(t *testing.T) {
	// Inicia o banco de dados
	dbStore, err := StartDatabase()
	if err != nil {
		t.Errorf("Falha ao criar banco de dados e tabelas: %s", err)
	}
	defer dbStore.DatabaseClose() // Garante que a conexão será fechada no final

	// storeUsers
	storeUsers := store.NewUser(dbStore.DB)
	// storeItensMenu
	storeItensMenu := store.NewMenuItem(dbStore.DB)
	// stoteItemPedido
	stoteItemPedido := store.NewItemOrder(dbStore.DB)
	// storePedido
	storePedido := store.NewOrder(dbStore.DB)

	// Cria usuários
	users, err := CreateUsers(storeUsers)
	if err != nil {
		t.Errorf("Falha ao criar banco de dados e tabelas: %s", err)
	}

	// Cria itens para o menu
	itens, err := CreateItensMenu(storeItensMenu)
	if err != nil {
		t.Errorf("Falha ao criar itensMenu: %s", err)
	}

	// Criar pedidos
	pedidos, err := CreateOrders(storePedido)
	if err != nil {
		t.Errorf("Falha ao criar pedidos: %s", err)
	}

	// Adicionar 1o itens ao pedido
	itemPedido := models.ItemOrder{
		OrderId:  pedidos[0].Id,
		ItemId:   itens[0].Id,
		Amount:   5,
		Comments: "Sem sal",
	}
	err = stoteItemPedido.CreateItemOrder(&itemPedido)
	if err != nil {
		t.Errorf("Falha ao adicionar item %s ao pedido %d: %s",
			itens[0].Name,
			pedidos[0].Id,
			err,
		)
	}

	// Adicionar 2o itens ao pedido
	itemPedido = models.ItemOrder{
		OrderId:  pedidos[0].Id,
		ItemId:   itens[1].Id,
		Amount:   2,
		Comments: "Mal passado",
	}
	err = stoteItemPedido.CreateItemOrder(&itemPedido)
	if err != nil {
		t.Errorf("Falha ao adicionar item %s ao pedido %d: %s",
			itens[0].Name,
			pedidos[0].Id,
			err,
		)
	}

	// Updade ItemPedido
	quantidade := itemPedido.Amount
	itemPedido.Amount += 1
	err = stoteItemPedido.UpdateItemOrder(&itemPedido)
	if err != nil {
		t.Errorf("Falha ao atualizar quantidade no itemPedido: %s", err)
	}

	// Carrega o ItemPedido
	itemPedidoLoad, err := stoteItemPedido.GetItemOrder(itemPedido.ItemId)
	if err != nil {
		t.Errorf("Erro ao ler um itemPedido: %s", err)
	}

	// Testa se quantidade é maior de 1
	if (quantidade + 1) != itemPedidoLoad.Amount {
		t.Errorf("Erro ao atualizar um itemPedido. Esperado %d, encontrado %d: %s",
			quantidade,
			itemPedidoLoad.Amount,
			err,
		)
	}

	// Pegar os pedidos do users[0]:
	allOrders0, err := storePedido.GetOrderByUser(users[0].Id)
	if err != nil {
		t.Errorf("Erro ao pegar os pedidos do usuário 0 um Pedido: %s", err)
	}
	if len(allOrders0) != 2 {
		t.Errorf("número de ordens experadas para o usuário %d eram 2, retornou %d",
			users[0].Id, len(allOrders0))
	}

	// Pegar os pedidos do users[1]:
	allOrders1, err := storePedido.GetOrderByUser(users[1].Id)
	if err != nil {
		t.Errorf("Erro ao pegar os pedidos do usuário 0 um Pedido: %s", err)
	}
	if len(allOrders1) != 1 {
		t.Errorf("número de ordens experadas para o usuário %d eram 1, retornou %d",
			users[1].Id, len(allOrders1))
	}

	// Pegar lista de pedidos pendentes
	pendingOrders, err := storePedido.GetPendingOrders()
	if err != nil {
		t.Errorf("Erro ao pegar os pedidos pendentes: %s", err)
	}
	if len(pendingOrders) != 4 {
		t.Errorf("esperado 4 pedidos pendentes. Encontrado %d", len(pendingOrders))
	}

	// Mudar status de pedidos[2]
	pedidos[2].Status = models.Entregue
	err = storePedido.UpdateOrder(&pedidos[2])
	if err != nil {
		t.Errorf("Erro ao atualizar um Pedido: %s", err)
	}
	// Pegar lista de pedidos pendentes
	pendingOrders, err = storePedido.GetPendingOrders()
	if err != nil {
		t.Errorf("Erro ao pegar os pedidos pendentes: %s", err)
	}
	if len(pendingOrders) != 3 {
		t.Errorf("esperado 3 pedidos pendentes. Encontrado %d", len(pendingOrders))
	}

	// Limpeza final
	if _, err := dbStore.DB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDBName)); err != nil {
		log.Fatal("Falha ao limpar o banco de dados de teste após o teste:", err)
	}
}
