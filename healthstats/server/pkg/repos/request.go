package repos

import (
	"healthstats/pkg/models"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type RequestRepo interface {
	CreateRequest(request models.Request) (string, error)
}

type requestRepo struct {
	tx     *sqlx.Tx
	logger zerolog.Logger
}

func NewRequestRepo(tx *sqlx.Tx, logger zerolog.Logger) RequestRepo {
	return &requestRepo{
		tx:     tx,
		logger: logger,
	}
}

func (r *requestRepo) CreateRequest(request models.Request) (string, error) {
	l := r.logger.With().Str("package", packageName).Str("function", "CreateRequest").Logger()

	query := `INSERT INTO requests (file_name, status) VALUES ($1, $2)`
	_, err := r.tx.Exec(query, request.FileName, request.Status)
	if err != nil {
		l.Err(err).Msg("failed to insert request")
		return "", err
	}

	return request.ID, nil
}
