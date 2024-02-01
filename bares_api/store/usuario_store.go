package store

import (
	"bares_api/models"
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

const (
  createUserSQL = "INSERT INTO %s(%s, %s, %s, %s) VALUES (?, ?, ?, ?)"
  getUserByEmailSQL = "SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?"
  updateUserSQL = "UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ? WHERE %s = ?"
  getUserSQL = "SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?"
  deleteUserSQL = "DELETE FROM %s WHERE %s = ?"
)

// UsuarioStore mantém a conexão com o banco de dados para operações relacionadas a usuários.
type UsuarioStore struct {
  DB *sql.DB
}

// NewUsuario cria uma nova instância de UsuarioStore.
func NewUsuario(db *sql.DB) *UsuarioStore {
  return &UsuarioStore{DB: db}
}

// UsuarioStorer define as operações que um UsuarioStore precisa implementar.
type UsuarioStorer interface {
  CreateUsuario(user *models.Usuario) error
  GetUsuario(id int) (*models.Usuario, error)
  GetUsuarioByEmail(email string) (*models.Usuario, error)
  UpdateUsuario(user *models.Usuario) error
  DeleteUsuario(id int) error
}

// Garanta que UsuarioStore implementa UsuarioStorer.
var _ UsuarioStorer = &UsuarioStore{}

// CreateUsuario adiciona um novo usuário ao banco de dados.
func (store *UsuarioStore) CreateUsuario(user *models.Usuario) error {
  sqlString := fmt.Sprintf(createUserSQL, TableUsuarios, Nome, Email, SenhaHash, Papel)

  stmt, err := store.DB.Prepare(sqlString)
  if err != nil {
    log.Printf("erro CreateUsuario: %v", err)
    return fmt.Errorf("erro CreateUsuario: %v", err)
  }
  defer stmt.Close()

  // hashed user password
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.SenhaHash), bcrypt.DefaultCost)
  if err != nil {
    log.Printf("erro CreateUsuario: %v", err)
    return fmt.Errorf("erro CreateUsuario: %v", err)
  }
  user.SenhaHash = string(hashedPassword)

  result, err := stmt.Exec(user.Nome, user.Email, user.SenhaHash, user.Papel)
  if err != nil {
    log.Printf("erro CreateUsuario: %v", err)
    return fmt.Errorf("erro CreateUsuario: %v", err)
  }

  usuarioID, err := result.LastInsertId()
  if err != nil {
    log.Printf("erro CreateUsuario: %v", err)
    return fmt.Errorf("erro CreateUsuario: %v", err)
  }
  user.UsuarioID = int(usuarioID)

  return nil
}

// GetUsuarioByEmail busca um usuário pelo Email.
func (store *UsuarioStore) GetUsuarioByEmail(email string) (*models.Usuario, error) {
  user := &models.Usuario{}

  sqlString := fmt.Sprintf(getUserByEmailSQL, UsuarioID, Nome, Email, SenhaHash, Papel, TableUsuarios, Email)

  err := store.DB.QueryRow(sqlString, email).Scan(
    &user.UsuarioID, &user.Nome, &user.Email, &user.SenhaHash, &user.Papel)
  if err != nil {
    log.Printf("erro GetUsuarioByEmail: %v", err)
    return nil, fmt.Errorf("erro GetUsuarioByEmail: %v", err)
  }

  return user, nil
}

// GetUsuario busca um usuário pelo ID.
func (store *UsuarioStore) GetUsuario(id int) (*models.Usuario, error) {
  user := &models.Usuario{}

  sqlString := fmt.Sprintf(getUserSQL, UsuarioID, Nome, Email, SenhaHash, Papel, TableUsuarios, UsuarioID)

  err := store.DB.QueryRow(sqlString, id).Scan(
    &user.UsuarioID,
    &user.Nome,
    &user.Email,
    &user.SenhaHash,
    &user.Papel,
  )

  if err != nil {
    log.Printf("erro GetUsuario: %v", err)
    return nil, fmt.Errorf("erro GetUsuario: %v", err)
  }
  return user, nil
}

// UpdateUsuario atualiza os dados de um usuário.
func (store *UsuarioStore) UpdateUsuario(user *models.Usuario) error {
  var hashedPassword string

  if user.SenhaHash != "" {
    // Hash da nova senha, se fornecida.
    var err error
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.SenhaHash), bcrypt.DefaultCost)
    if err != nil {
      log.Printf("erro UpdateUsuario: %v", err)
      return fmt.Errorf("erro UpdateUsuario: %v", err)
    }
    hashedPassword = string(hashedBytes)
  } else {
    // Recupera a senha atual (hash) para não substituir por vazio.
    currentUser, err := store.GetUsuario(user.UsuarioID)
    if err != nil {
      log.Printf("erro UpdateUsuario: %v", err)
      return fmt.Errorf("erro UpdateUsuario: %v", err)
    }
    hashedPassword = currentUser.SenhaHash
  }

  sqlString := fmt.Sprintf(updateUserSQL, TableUsuarios, Nome, Email, SenhaHash, Papel, UsuarioID)
  stmt, err := store.DB.Prepare(sqlString)
  if err != nil {
    log.Printf("erro UpdateUsuario: %v", err)
    return fmt.Errorf("erro UpdateUsuario: %v", err)
  }
  defer stmt.Close()

  _, err = stmt.Exec(user.Nome, user.Email, hashedPassword, user.Papel, user.UsuarioID)
  if err != nil {
    log.Printf("erro UpdateUsuario: %v", err)
    return fmt.Errorf("erro UpdateUsuario: %v", err)
  }

  return nil
}

// DeleteUsuario remove um usuário do banco de dados.
// FIXME: as remoções de registros das tabelas do banco de dados devem ser tratadas
// com cuidado, que não serão tomados aqui pelo carater de estudo este código.
func (store *UsuarioStore) DeleteUsuario(id int) error {
  sqlString := fmt.Sprintf(deleteUserSQL, TableUsuarios, UsuarioID)
  stmt, err := store.DB.Prepare(sqlString)
  if err != nil {
    log.Printf("erro DeleteUsuario: %v", err)
    return fmt.Errorf("erro DeleteUsuario: %v", err)
  }
  defer stmt.Close()

  _, err = stmt.Exec(id)
  if err != nil {
    log.Printf("erro DeleteUsuario: %v", err)
    return fmt.Errorf("erro DeleteUsuario: %v", err)
  }

  return nil
}
