package integration

import (
	"bares_api/models"
	"bares_api/store"
	"fmt"
	"log"
	"testing"
)

func TestPedidoIntegration(t *testing.T) {
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
	pedidos, err := CreatePedidos(storePedido)
	if err != nil {
		t.Errorf("Falha ao criar pedidos: %s", err)
	}

	// Adicionar 1o itens ao pedido
	itemPedido := models.ItemOrder{
		PedidoID:    pedidos[0].PedidoID,
		ItemID:      itens[0].ItemID,
		Quantidade:  5,
		Observacoes: "Sem sal",
	}
	err = stoteItemPedido.CreateItemOrder(&itemPedido)
	if err != nil {
		t.Errorf("Falha ao adicionar item %s ao pedido %d: %s",
			itens[0].Nome,
			pedidos[0].PedidoID,
			err,
		)
	}

	// Adicionar 2o itens ao pedido
	itemPedido = models.ItemOrder{
		PedidoID:    pedidos[0].PedidoID,
		ItemID:      itens[1].ItemID,
		Quantidade:  2,
		Observacoes: "Mal passado",
	}
	err = stoteItemPedido.CreateItemOrder(&itemPedido)
	if err != nil {
		t.Errorf("Falha ao adicionar item %s ao pedido %d: %s",
			itens[0].Nome,
			pedidos[0].PedidoID,
			err,
		)
	}

	// Updade ItemPedido
	quantidade := itemPedido.Quantidade
	itemPedido.Quantidade += 1
	err = stoteItemPedido.UpdateItemOrder(&itemPedido)
	if err != nil {
		t.Errorf("Falha ao atualizar quantidade no itemPedido: %s", err)
	}

	// Carrega o ItemPedido
	itemPedidoLoad, err := stoteItemPedido.GetItemOrder(itemPedido.ItemID)
	if err != nil {
		t.Errorf("Erro ao ler um itemPedido: %s", err)
	}

	// Testa se quantidade é maior de 1
	if (quantidade + 1) != itemPedidoLoad.Quantidade {
		t.Errorf("Erro ao atualizar um itemPedido. Esperado %d, encontrado %d: %s",
			quantidade,
			itemPedidoLoad.Quantidade,
			err,
		)
	}

	// Pegar os pedidos do users[0]:
	allOrders0, err := storePedido.GetOrderByUser(users[0].UsuarioID)
	if err != nil {
		t.Errorf("Erro ao pegar os pedidos do usuário 0 um Pedido: %s", err)
	}
	if len(allOrders0) != 2 {
		t.Errorf("número de ordens experadas para o usuário %d eram 2, retornou %d",
			users[0].UsuarioID, len(allOrders0))
	}

	// Pegar os pedidos do users[1]:
	allOrders1, err := storePedido.GetOrderByUser(users[1].UsuarioID)
	if err != nil {
		t.Errorf("Erro ao pegar os pedidos do usuário 0 um Pedido: %s", err)
	}
	if len(allOrders1) != 1 {
		t.Errorf("número de ordens experadas para o usuário %d eram 1, retornou %d",
			users[1].UsuarioID, len(allOrders1))
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
