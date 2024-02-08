package storetest_test

import (
	"bares_api/models"
	"bares_api/store"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUsuario(t *testing.T) {
	// Cria um mock do banco de dados e do sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Cria um mock do resultado do insert
	mockResult := sqlmock.NewResult(1, 1)

	// Prepara o mock para esperar a chamada que será feita pela função CreateUsuario
	mock.ExpectPrepare("INSERT INTO UsersTable").
		ExpectExec().
		WithArgs(
			"John Doe",
			"johndoe@example.com",
			sqlmock.AnyArg(),
			"cliente",
		).
		WillReturnResult(mockResult)

	// Cria uma instância de UsuarioStore com o mock do banco de dados
	store := store.NewUser(db)

	// Cria um usuário de teste
	testUser := &models.User{
		Name:         "John Doe",
		Email:        "johndoe@example.com",
		PasswordHash: "senha123",
		Role:         "cliente",
	}

	// Chama a função CreateUsuario
	err = store.CreateUser(testUser)
	assert.NoError(t, err)

	// Verifica se todas as expectativas foram atendidas
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUsuarioByEmail(t *testing.T) {
	// Cria um mock do banco de dados e do sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Define as colunas que serão retornadas pelo mock do banco de dados
	columns := []string{"usuarioID", "nome", "email", "senhaHash", "papel"}

	// Define o resultado esperado do mock
	expectedResult := sqlmock.NewRows(columns).
		AddRow(
			1,
			"John Doe",
			"johndoe@example.com",
			"hashedpassword",
			"cliente",
		)

	// Prepara o mock para esperar a chamada que será feita pela função GetUsuarioByEmail
	mock.ExpectQuery(
		"SELECT id, name, email, passwordHash, role FROM UsersTable WHERE email = ?").
		WithArgs("johndoe@example.com").
		WillReturnRows(expectedResult)

	// Cria uma instância de UsuarioStore com o mock do banco de dados
	store := store.NewUser(db)

	// Chama a função GetUsuarioByEmail
	user, err := store.GetUserByEmail("johndoe@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "johndoe@example.com", user.Email)
	assert.Equal(t, "hashedpassword", user.PasswordHash)
	assert.Equal(t, "cliente", string(user.Role))

	// Verifica se todas as expectativas foram atendidas
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUsuario(t *testing.T) {
	// Cria um mock do banco de dados e do sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Define as colunas que serão retornadas pelo mock do banco de dados
	columns := []string{"id", "name", "email", "passwordHash", "role"}
	// Define o resultado esperado do mock
	expectedResult := sqlmock.
		NewRows(columns).
		AddRow(
			1,
			"John Doe",
			"johndoe@example.com",
			"hashedpassword",
			"cliente",
		)

	// Prepara o mock para esperar a chamada que será feita pela função GetUsuario
	mock.ExpectQuery(
		"SELECT id, name, email, passwordHash, role FROM UsersTable WHERE id = ?").
		WithArgs(1).
		WillReturnRows(expectedResult)

	// Cria uma instância de UsuarioStore com o mock do banco de dados
	store := store.NewUser(db)

	// Chama a função GetUsuario
	user, err := store.GetUser(1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.Id)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "johndoe@example.com", user.Email)
	assert.Equal(t, "hashedpassword", user.PasswordHash)
	assert.Equal(t, "cliente", string(user.Role))

	// Verifica se todas as expectativas foram atendidas
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateUsuario(t *testing.T) {
	// Cria um mock do banco de dados e do sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Cria um mock do resultado da operação de atualização
	mockResult := sqlmock.NewResult(0, 1) // 0 LastInsertId, 1 RowsAffected

	// Prepara o mock para esperar a chamada que será feita pela função UpdateUsuario
	mock.ExpectPrepare("UPDATE UsersTable SET").
		ExpectExec().
		WithArgs(
			"John Doe Updated",
			"johndoeupdated@example.com",
			sqlmock.AnyArg(),
			"cliente",
			1,
		).
		WillReturnResult(mockResult)

	// Cria uma instância de UsuarioStore com o mock do banco de dados
	store := store.NewUser(db)

	// Cria um usuário de teste
	testUser := &models.User{
		Id:           1,
		Name:         "John Doe Updated",
		Email:        "johndoeupdated@example.com",
		PasswordHash: "newhashedpassword",
		Role:         "cliente",
	}

	// Chama a função UpdateUsuario
	err = store.UpdateUser(testUser)
	assert.NoError(t, err)

	// Verifica se todas as expectativas foram atendidas
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteUsuario(t *testing.T) {
	// Cria um mock do banco de dados e do sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Cria um mock do resultado da operação de exclusão
	mockResult := sqlmock.NewResult(0, 1) // 0 LastInsertId, 1 RowsAffected

	// Prepara o mock para esperar a chamada que será feita pela função DeleteUsuario
	mock.ExpectPrepare("DELETE FROM UsersTable").
		ExpectExec().
		WithArgs(1).
		WillReturnResult(mockResult)

	// Cria uma instância de UsuarioStore com o mock do banco de dados
	store := store.NewUser(db)

	// Chama a função DeleteUsuario
	err = store.DeleteUser(1)
	assert.NoError(t, err)

	// Verifica se todas as expectativas foram atendidas
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
