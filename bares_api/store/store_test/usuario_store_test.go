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
  mock.ExpectPrepare("INSERT INTO Usuarios").
    ExpectExec().
    WithArgs(
      "John Doe",
      "johndoe@example.com",
      "hashedpassword",
      "cliente",
    ).
    WillReturnResult(mockResult)

  // Cria uma instância de UsuarioStore com o mock do banco de dados
  store := store.NewUsuario(db)

  // Cria um usuário de teste
  testUser := &models.Usuario{
    Nome:      "John Doe",
    Email:     "johndoe@example.com",
    SenhaHash: "hashedpassword",
    Papel:     "cliente",
  }

  // Chama a função CreateUsuario
  err = store.CreateUsuario(testUser)
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
    "SELECT usuarioID, nome, email, senhaHash, papel FROM Usuarios WHERE email = ?").
    WithArgs("johndoe@example.com").
    WillReturnRows(expectedResult)

  // Cria uma instância de UsuarioStore com o mock do banco de dados
  store := store.NewUsuario(db)

  // Chama a função GetUsuarioByEmail
  user, err := store.GetUsuarioByEmail("johndoe@example.com")
  assert.NoError(t, err)
  assert.NotNil(t, user)
  assert.Equal(t, "John Doe", user.Nome)
  assert.Equal(t, "johndoe@example.com", user.Email)
  assert.Equal(t, "hashedpassword", user.SenhaHash)
  assert.Equal(t, "cliente", string(user.Papel))

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
  columns := []string{"usuarioID", "nome", "email", "senhaHash", "papel"}
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
    "SELECT usuarioID, nome, email, senhaHash, papel FROM Usuarios WHERE usuarioID = ?").
    WithArgs(1).
    WillReturnRows(expectedResult)

  // Cria uma instância de UsuarioStore com o mock do banco de dados
  store := store.NewUsuario(db)

  // Chama a função GetUsuario
  user, err := store.GetUsuario(1)
  assert.NoError(t, err)
  assert.NotNil(t, user)
  assert.Equal(t, 1, user.UsuarioID)
  assert.Equal(t, "John Doe", user.Nome)
  assert.Equal(t, "johndoe@example.com", user.Email)
  assert.Equal(t, "hashedpassword", user.SenhaHash)
  assert.Equal(t, "cliente", string(user.Papel))

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
  mock.ExpectPrepare("UPDATE Usuarios SET").
    ExpectExec().
    WithArgs(
      "John Doe Updated",
      "johndoeupdated@example.com",
      "newhashedpassword",
      "cliente",
      1,
    ).
    WillReturnResult(mockResult)

  // Cria uma instância de UsuarioStore com o mock do banco de dados
  store := store.NewUsuario(db)

  // Cria um usuário de teste
  testUser := &models.Usuario{
    UsuarioID: 1,
    Nome:      "John Doe Updated",
    Email:     "johndoeupdated@example.com",
    SenhaHash: "newhashedpassword",
    Papel:     "cliente",
  }

  // Chama a função UpdateUsuario
  err = store.UpdateUsuario(testUser)
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
  mock.ExpectPrepare("DELETE FROM Usuarios").
    ExpectExec().
    WithArgs(1).
    WillReturnResult(mockResult)

  // Cria uma instância de UsuarioStore com o mock do banco de dados
  store := store.NewUsuario(db)

  // Chama a função DeleteUsuario
  err = store.DeleteUsuario(1)
  assert.NoError(t, err)

  // Verifica se todas as expectativas foram atendidas
  if err := mock.ExpectationsWereMet(); err != nil {
    t.Errorf("there were unfulfilled expectations: %s", err)
  }
}
