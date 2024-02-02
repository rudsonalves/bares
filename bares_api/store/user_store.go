package store

import (
	"bares_api/models"
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

const (
	createUserSQL     = "INSERT INTO %s(%s, %s, %s, %s) VALUES (?, ?, ?, ?)"
	getUserByEmailSQL = "SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?"
	updateUserSQL     = "UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ? WHERE %s = ?"
	getUserSQL        = "SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?"
	deleteUserSQL     = "DELETE FROM %s WHERE %s = ?"
)

// UserStore mantém a conexão com o banco de dados para operações relacionadas a usuários.
type UserStore struct {
	DB *sql.DB
}

// NewUser cria uma nova instância de UsuarioStore.
func NewUser(db *sql.DB) *UserStore {
	return &UserStore{DB: db}
}

// UsuarioStorer define as operações que um UsuarioStore precisa implementar.
type UsuarioStorer interface {
	CreateUser(user *models.User) error
	GetUser(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

// Garanta que UsuarioStore implementa UsuarioStorer.
var _ UsuarioStorer = &UserStore{}

// CreateUser adiciona um novo usuário ao banco de dados.
func (store *UserStore) CreateUser(user *models.User) error {
	sqlString := fmt.Sprintf(createUserSQL, TableUsers, Name, Email, PasswordHash, Role)

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

// GetUserByEmail busca um usuário pelo Email.
func (store *UserStore) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	sqlString := fmt.Sprintf(getUserByEmailSQL, UserID, Name, Email, PasswordHash, Role, TableUsers, Email)

	err := store.DB.QueryRow(sqlString, email).Scan(
		&user.UsuarioID, &user.Nome, &user.Email, &user.SenhaHash, &user.Papel)
	if err != nil {
		log.Printf("erro GetUsuarioByEmail: %v", err)
		return nil, fmt.Errorf("erro GetUsuarioByEmail: %v", err)
	}

	return user, nil
}

// GetUser busca um usuário pelo ID.
func (store *UserStore) GetUser(id int) (*models.User, error) {
	user := &models.User{}

	sqlString := fmt.Sprintf(getUserSQL, UserID, Name, Email, PasswordHash, Role, TableUsers, UserID)

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

// UpdateUser atualiza os dados de um usuário.
func (store *UserStore) UpdateUser(user *models.User) error {
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
		currentUser, err := store.GetUser(user.UsuarioID)
		if err != nil {
			log.Printf("erro UpdateUsuario: %v", err)
			return fmt.Errorf("erro UpdateUsuario: %v", err)
		}
		hashedPassword = currentUser.SenhaHash
	}

	sqlString := fmt.Sprintf(updateUserSQL, TableUsers, Name, Email, PasswordHash, Role, UserID)
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

// DeleteUser remove um usuário do banco de dados.
// FIXME: as remoções de registros das tabelas do banco de dados devem ser tratadas
// com cuidado, que não serão tomados aqui pelo carater de estudo este código.
func (store *UserStore) DeleteUser(id int) error {
	sqlString := fmt.Sprintf(deleteUserSQL, TableUsers, UserID)
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
