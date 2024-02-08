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
	mock.ExpectPrepare("INSERT INTO MenuItemsTable").
		ExpectExec().
		WithArgs("Pizza", "Delicious pizza", 9.99, "imageurl").
		WillReturnResult(mockResult)

	store := store.NewMenuItem(db)
	testItemMenu := &models.MenuItem{
		Name:        "Pizza",
		Description: "Delicious pizza",
		Price:       9.99,
		ImageURL:    "imageurl",
	}

	err = store.CreateMenuItem(testItemMenu)
	assert.NoError(t, err)
	assert.Equal(t, 1, testItemMenu.Id)

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

	columns := []string{"id", "name", "description", "price", "imageURL"}
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
		"SELECT id, name, description, price, imageURL FROM MenuItemsTable WHERE id = ?").
		WithArgs(1).
		WillReturnRows(expectedResult)

	store := store.NewMenuItem(db)

	result, err := store.GetMenuItem(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "Pizza", result.Name)
	assert.Equal(t, "Delicious pizza", result.Description)
	assert.Equal(t, 9.99, result.Price)
	assert.Equal(t, "imageurl", result.ImageURL)

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
	mock.ExpectPrepare("UPDATE MenuItemsTable SET").
		ExpectExec().
		WithArgs(
			"Pizza Updated",
			"Delicious pizza updated",
			10.99,
			"imageurlupdated",
			1,
		).
		WillReturnResult(mockResult)

	store := store.NewMenuItem(db)
	testItemMenu := &models.MenuItem{
		Id:          1,
		Name:        "Pizza Updated",
		Description: "Delicious pizza updated",
		Price:       10.99,
		ImageURL:    "imageurlupdated",
	}

	err = store.UpdateMenuItem(testItemMenu)
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
	mock.ExpectPrepare("DELETE FROM MenuItemsTable").
		ExpectExec().
		WithArgs(1).
		WillReturnResult(mockResult)

	store := store.NewMenuItem(db)

	err = store.DeleteMenuItem(1)
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

	columns := []string{"id", "name", "description", "price", "imageURL"}
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
		"SELECT id, name, description, price, imageURL FROM MenuItemsTable ORDER BY name").
		WillReturnRows(expectedResult)

	store := store.NewMenuItem(db)

	result, err := store.GetAllMenuItem()
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

	columns := []string{"id", "name", "description", "price", "imageURL"}
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
		"SELECT id, name, description, price, imageURL FROM MenuItemsTable WHERE name = ?").
		WithArgs("Pizza").
		WillReturnRows(expectedResult)

	store := store.NewMenuItem(db)

	result, err := store.GetMenuItemByName("Pizza")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "Pizza", result.Name)
	assert.Equal(t, "Delicious pizza", result.Description)
	assert.Equal(t, 9.99, result.Price)
	assert.Equal(t, "imageurl", result.ImageURL)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
