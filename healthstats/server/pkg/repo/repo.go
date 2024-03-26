package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

const (
	packageName = "repo"
)

type Repository struct {
	tx          *sqlx.Tx
	logger      zerolog.Logger
	RequestRepo RequestRepo
}

func NewRepository(tx *sqlx.Tx, logger zerolog.Logger) *Repository {

	requestRepo := NewRequestRepo(tx, logger)

	return &Repository{
		tx:          tx,
		logger:      logger,
		RequestRepo: requestRepo,
	}
}
