package config

import (
	"database/sql"
	"fmt"

	"github.com/ghulammuzz/go-restful-template/pkg/env"
	"github.com/ghulammuzz/go-restful-template/pkg/logger"
	_ "github.com/lib/pq"
)

func Initialize() (*sql.DB, error) {

	dsn := env.GetEnv("DATABASE_URL")
	logger.Debug("database url %s : ", dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping the database: %v", err)
	}

	return db, nil
}
