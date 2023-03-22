package repositories

import (
	"api/entitites"
	"database/sql"
	"log"
)

type UserRepository struct {
	connection *sql.DB
}

func InitUserRepository(connection *sql.DB) *UserRepository {
	return &UserRepository{connection: connection}
}

type IUserRepository interface {
	Get() (*[]entitites.User, error)
	GetByID(id uint) (*entitites.User, error)
	Create(payload *entitites.User) error
	Update(id uint, payload *entitites.User) error
	Delete(id uint) error
}

func (repo *UserRepository) Get() (*[]entitites.User, error) {
	var (
		users = []entitites.User{}
	)

	rows, err := repo.connection.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u entitites.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}

func (repo *UserRepository) GetByID(id uint) (*entitites.User, error) {
	var u entitites.User

	err := repo.connection.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (repo *UserRepository) Create(payload *entitites.User) error {
	err := repo.connection.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", payload.Name, payload.Email).Scan(&payload.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) Update(id uint, payload *entitites.User) error {
	_, err := repo.connection.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", payload.Name, payload.Email, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Delete(id uint) error {
	_, err := repo.connection.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
