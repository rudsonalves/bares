package store

import (
	"bares_api/models"
	"database/sql"
	"fmt"
)

// UsuarioStore mantém a conexão com o banco de dados para operações relacionadas a usuários.
type UsuarioStore struct {
	DB *sql.DB
}

// NewUsuarioStore cria uma nova instância de UsuarioStore.
func NewUsuarioStore(db *sql.DB) *UsuarioStore {
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
	createString := fmt.Sprintf("INSERT INTO %s(%s, %s, %s, %s) VALUES (?, ?, ?, ?)",
		TableUsuarios, Nome, Email, SenhaHash, Papel)

	stmt, err := store.DB.Prepare(createString)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Nome, user.Email, user.SenhaHash, user.Papel)
	if err != nil {
		return err
	}

	usuarioID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.UsuarioID = int(usuarioID)

	return nil
}

// GetUsuarioByEmail busca um usuário pelo Email.
func (store *UsuarioStore) GetUsuarioByEmail(email string) (*models.Usuario, error) {
	user := &models.Usuario{}

	queryString := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		UsuarioID, Nome, Email, SenhaHash, Papel, TableUsuarios, Email)

	err := store.DB.QueryRow(queryString, email).Scan(
		&user.UsuarioID, &user.Nome, &user.Email, &user.SenhaHash, &user.Papel)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUsuario busca um usuário pelo ID.
func (store *UsuarioStore) GetUsuario(id int) (*models.Usuario, error) {
	u := &models.Usuario{}

	queryString := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		UsuarioID, Nome, Email, SenhaHash, Papel, TableUsuarios, UsuarioID)

	err := store.DB.QueryRow(queryString, id).Scan(&u.UsuarioID, &u.Nome, &u.Email, &u.SenhaHash, &u.Papel)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateUsuario atualiza os dados de um usuário.
func (store *UsuarioStore) UpdateUsuario(user *models.Usuario) error {
	updateString := fmt.Sprintf("UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ? WHERE %s = ?",
		TableUsuarios, Nome, Email, SenhaHash, Papel, UsuarioID)

	stmt, err := store.DB.Prepare(updateString)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Nome, user.Email, user.SenhaHash, user.Papel, user.UsuarioID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUsuario remove um usuário do banco de dados.
func (store *UsuarioStore) DeleteUsuario(id int) error {
	deleteString := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", TableUsuarios, UsuarioID)
	stmt, err := store.DB.Prepare(deleteString)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
