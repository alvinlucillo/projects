package model

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

const (
	packageName = "models"
)

type PostgreSQLDBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewPostgreSQLDB(config PostgreSQLDBConfig, logger zerolog.Logger) (*sqlx.DB, error) {
	l := logger.With().Str("package", packageName).Str("function", "NewPostgreSQLDB").Logger()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.DBName)

	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		l.Err(err).Msg("failed to open db")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		l.Err(err).Msg("failed to ping db")
		return nil, err
	}
	return db, nil
}
