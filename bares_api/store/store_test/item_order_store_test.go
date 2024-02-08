package storetest_test

import (
	"bares_api/models"
	"bares_api/store"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateItemPedidoStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockResult := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("INSERT INTO ItemsOrderTable").
		ExpectExec().
		WithArgs(1, 1, 10, "Sem cebola").
		WillReturnResult(mockResult)

	store := store.NewItemOrder(db)
	testItemPedido := &models.ItemOrder{
		OrderId:  1,
		ItemId:   1,
		Amount:   10,
		Comments: "Sem cebola",
	}

	err = store.CreateItemOrder(testItemPedido)
	assert.NoError(t, err)
	assert.Equal(t, 1, testItemPedido.Id)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetItemPedidoStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	columns := []string{"id", "orderId", "itemId", "amount", "comments"}
	expectedResult := sqlmock.NewRows(columns).AddRow(1, 1, 1, 10, "Sem cebola")
	mock.ExpectQuery("SELECT id, orderId, itemId, amount, comments FROM ItemsOrderTable WHERE id = ?").
		WithArgs(1).
		WillReturnRows(expectedResult)

	store := store.NewItemOrder(db)

	result, err := store.GetItemOrder(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, 1, result.OrderId)
	assert.Equal(t, 1, result.ItemId)
	assert.Equal(t, 10, result.Amount)
	assert.Equal(t, "Sem cebola", result.Comments)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateItemPedidoStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockResult := sqlmock.NewResult(0, 1)
	mock.ExpectPrepare("UPDATE ItemsOrderTable SET").
		ExpectExec().
		WithArgs(1, 1, 15, "Adicionar molho extra", 1).
		WillReturnResult(mockResult)

	store := store.NewItemOrder(db)
	testItemPedido := &models.ItemOrder{
		Id:       1,
		OrderId:  1,
		ItemId:   1,
		Amount:   15,
		Comments: "Adicionar molho extra",
	}

	err = store.UpdateItemOrder(testItemPedido)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteItemPedidoStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockResult := sqlmock.NewResult(0, 1)
	mock.ExpectPrepare("DELETE FROM ItemsOrderTable").
		ExpectExec().
		WithArgs(1).
		WillReturnResult(mockResult)

	store := store.NewItemOrder(db)

	err = store.DeleteItemOrder(1)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
