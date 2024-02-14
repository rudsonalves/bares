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
	getUsersByRole    = "SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?"
	getAllUsersSQL    = "SELECT %s, %s, %s, %s, %s FROM %s"
	updateUserSQL     = "UPDATE %s SET %s = ?, %s = ?, %s = ? WHERE %s = ?"
	updateUserPassSQL = "UPDATE %s SET %s = ? WHERE %s = ?"
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
	UpdateUserPass(userId int, password string) error
	DeleteUser(id int) error
	GetUsersByRole(role models.Role) ([]*models.User, error)
	GetAllUsers() ([]*models.User, error)
}

// Garanta que UsuarioStore implementa UsuarioStorer.
var _ UsuarioStorer = &UserStore{}

// CreateUser adiciona um novo usuário ao banco de dados.
func (store *UserStore) CreateUser(user *models.User) error {
	sqlString := fmt.Sprintf(createUserSQL, TableUsers, Name, Email, PasswordHash, Role)

	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro CreateUsuario: %v", err)
		return err
	}
	defer stmt.Close()

	// hashed user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("erro CreateUsuario: %v", err)
		return err
	}
	user.PasswordHash = string(hashedPassword)

	result, err := stmt.Exec(user.Name, user.Email, user.PasswordHash, user.Role)
	if err != nil {
		log.Printf("erro CreateUsuario: %v", err)
		return err
	}

	usuarioID, err := result.LastInsertId()
	if err != nil {
		log.Printf("erro CreateUsuario: %v", err)
		return err
	}
	user.Id = int(usuarioID)

	return nil
}

// GetUserByEmail busca um usuário pelo Email.
func (store *UserStore) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	sqlString := fmt.Sprintf(getUserByEmailSQL, Id, Name, Email, PasswordHash, Role, TableUsers, Email)

	err := store.DB.QueryRow(sqlString, email).Scan(
		&user.Id, &user.Name, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		log.Printf("erro GetUsuarioByEmail: %v", err)
		return nil, err
	}

	return user, nil
}

// GetUser busca um usuário pelo ID.
func (store *UserStore) GetUser(id int) (*models.User, error) {
	user := &models.User{}

	sqlString := fmt.Sprintf(getUserSQL, Id, Name, Email, PasswordHash, Role, TableUsers, Id)

	err := store.DB.QueryRow(sqlString, id).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
	)

	if err != nil {
		log.Printf("erro GetUsuario: %v", err)
		return nil, err
	}
	return user, nil
}

// GetUserByRole busca um usuário pelo seu role
func (store *UserStore) GetUsersByRole(role models.Role) ([]*models.User, error) {
	sqlString := fmt.Sprintf(getUsersByRole, Id, Name, Email, PasswordHash, Role,
		TableUsers, Role)

	rows, err := store.DB.Query(sqlString, role)
	if err != nil {
		log.Printf("erro GetUsersByRole: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.PasswordHash, &user.Role); err != nil {
			log.Printf("erro GetUsersByRole: %v", err)
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		log.Printf("erro GetUsersByRole: %v", err)
		return nil, err
	}

	return users, nil
}

func (store *UserStore) GetAllUsers() ([]*models.User, error) {
	sqlString := fmt.Sprintf(getAllUsersSQL, Id, Name, Email, PasswordHash, Role, TableUsers)

	rows, err := store.DB.Query(sqlString)
	if err != nil {
		log.Printf("error GetAllUsers: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.PasswordHash, &user.Role); err != nil {
			log.Printf("erro GetAllUsers: %v", err)
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		log.Printf("erro GetAllUsers: %v", err)
		return nil, err
	}

	return users, nil
}

// UpdateUser updates a user's password.
func (store *UserStore) UpdateUserPass(userId int, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("erro UpdateUserPass: %v", err)
		return err
	}

	sqlString := fmt.Sprintf(updateUserPassSQL, TableUsers, PasswordHash, Id)
	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro UpdateUserPass: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(hashedPassword, userId)
	if err != nil {
		log.Printf("erro UpdateUserPass: %v", err)
		return err
	}

	return nil
}

// UpdateUser updates a user's data, except the password.
func (store *UserStore) UpdateUser(user *models.User) error {
	// var hashedPassword string
	//
	// if user.PasswordHash != "" {
	// 	// Hash da nova senha, se fornecida.
	// 	var err error
	// 	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	// 	if err != nil {
	// 		log.Printf("erro UpdateUser: %v", err)
	// 		return err
	// 	}
	// 	hashedPassword = string(hashedBytes)
	// } else {
	// 	// Recupera a senha atual (hash) para não substituir por vazio.
	// 	currentUser, err := store.GetUser(user.Id)
	// 	if err != nil {
	// 		log.Printf("erro UpdateUser: %v", err)
	// 		return err
	// 	}
	// 	hashedPassword = currentUser.PasswordHash
	// }

	sqlString := fmt.Sprintf(updateUserSQL, TableUsers, Name, Email, Role, Id)
	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro UpdateUser: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Role, user.Id)
	if err != nil {
		log.Printf("erro UpdateUser: %v", err)
		return err
	}

	return nil
}

// DeleteUser remove um usuário do banco de dados.
// FIXME: as remoções de registros das tabelas do banco de dados devem ser tratadas
// com cuidado, que não serão tomados aqui pelo carater de estudo este código.
func (store *UserStore) DeleteUser(id int) error {
	sqlString := fmt.Sprintf(deleteUserSQL, TableUsers, Id)
	stmt, err := store.DB.Prepare(sqlString)
	if err != nil {
		log.Printf("erro DeleteUsuario: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Printf("erro DeleteUsuario: %v", err)
		return err
	}

	return nil
}
