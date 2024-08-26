package repository

import (
	"database/sql"

	"github.com/ghulammuzz/go-restful-template/internal/model"
	_ "github.com/lib/pq"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(user *model.User) error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err := r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}
