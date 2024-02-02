package storetest_test

import (
	"bares_api/models"
	"bares_api/store"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateItemMenu(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockResult := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("INSERT INTO ItensMenu").
		ExpectExec().
		WithArgs("Pizza", "Delicious pizza", 9.99, "imageurl").
		WillReturnResult(mockResult)

	store := store.NewItensMenu(db)
	testItemMenu := &models.ItemMenu{
		Nome:      "Pizza",
		Descricao: "Delicious pizza",
		Preco:     9.99,
		ImagemURL: "imageurl",
	}

	err = store.CreateItemMenu(testItemMenu)
	assert.NoError(t, err)
	assert.Equal(t, 1, testItemMenu.ItemID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetItemMenu(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	columns := []string{"itemID", "nome", "descricao", "preco", "imagemURL"}
	expectedResult := sqlmock.
		NewRows(columns).
		AddRow(
			1,
			"Pizza",
			"Delicious pizza",
			9.99,
			"imageurl",
		)
	mock.ExpectQuery(
		"SELECT itemID, nome, descricao, preco, imagemURL FROM ItensMenu WHERE itemID = ?").
		WithArgs(1).
		WillReturnRows(expectedResult)

	store := store.NewItensMenu(db)

	result, err := store.GetItemMenu(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ItemID)
	assert.Equal(t, "Pizza", result.Nome)
	assert.Equal(t, "Delicious pizza", result.Descricao)
	assert.Equal(t, 9.99, result.Preco)
	assert.Equal(t, "imageurl", result.ImagemURL)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateItemMenu(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockResult := sqlmock.NewResult(0, 1)
	mock.ExpectPrepare("UPDATE ItensMenu SET").
		ExpectExec().
		WithArgs(
			"Pizza Updated",
			"Delicious pizza updated",
			10.99,
			"imageurlupdated",
			1,
		).
		WillReturnResult(mockResult)

	store := store.NewItensMenu(db)
	testItemMenu := &models.ItemMenu{
		ItemID:    1,
		Nome:      "Pizza Updated",
		Descricao: "Delicious pizza updated",
		Preco:     10.99,
		ImagemURL: "imageurlupdated",
	}

	err = store.UpdateItemMenu(testItemMenu)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteItemMenu(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockResult := sqlmock.NewResult(0, 1)
	mock.ExpectPrepare("DELETE FROM ItensMenu").
		ExpectExec().
		WithArgs(1).
		WillReturnResult(mockResult)

	store := store.NewItensMenu(db)

	err = store.DeleteItemMenu(1)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllItemMenu(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	columns := []string{"itemID", "nome", "descricao", "preco", "imagemURL"}
	expectedResult := sqlmock.NewRows(columns).
		AddRow(
			1,
			"Pizza",
			"Delicious pizza",
			9.99,
			"imageurl",
		).
		AddRow(
			2,
			"Burger",
			"Tasty burger",
			7.99,
			"imageurl2",
		)
	mock.ExpectQuery(
		"SELECT itemID, nome, descricao, preco, imagemURL FROM ItensMenu ORDER BY nome").
		WillReturnRows(expectedResult)

	store := store.NewItensMenu(db)

	result, err := store.GetAllItemMenu()
	assert.NoError(t, err)
	assert.Len(t, result, 2)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetItemMenuByNome(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	columns := []string{"itemID", "nome", "descricao", "preco", "imagemURL"}
	expectedResult := sqlmock.
		NewRows(columns).
		AddRow(
			1,
			"Pizza",
			"Delicious pizza",
			9.99,
			"imageurl",
		)
	mock.ExpectQuery(
		"SELECT itemID, nome, descricao, preco, imagemURL FROM ItensMenu WHERE nome = ?").
		WithArgs("Pizza").
		WillReturnRows(expectedResult)

	store := store.NewItensMenu(db)

	result, err := store.GetItemMenuByNome("Pizza")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ItemID)
	assert.Equal(t, "Pizza", result.Nome)
	assert.Equal(t, "Delicious pizza", result.Descricao)
	assert.Equal(t, 9.99, result.Preco)
	assert.Equal(t, "imageurl", result.ImagemURL)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
