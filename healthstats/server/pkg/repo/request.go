package repo

import (
	"healthstats/pkg/model"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type RequestRepo interface {
	CreateRequest(request model.Request) (string, error)
	UpdateRequest(request model.Request) error
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

func (r *requestRepo) CreateRequest(request model.Request) (string, error) {
	l := r.logger.With().Str("package", packageName).Str("function", "CreateRequest").Logger()

	query := `INSERT INTO request (file_name, status) VALUES ($1, $2) RETURNING id;`

	var id string
	err := r.tx.QueryRow(query, request.FileName, request.Status).Scan(&id)
	if err != nil {
		l.Err(err).Msg("failed to insert request")
		return "", err
	}

	return id, nil
}

func (r *requestRepo) UpdateRequest(request model.Request) error {
	l := r.logger.With().Str("package", packageName).Str("function", "UpdateRequest").Logger()

	query := `UPDATE request SET status = $1, file_name = $2, updated_at = NOW() WHERE id = $3;`
	_, err := r.tx.Exec(query, request.Status, request.FileName, request.ID)
	if err != nil {
		l.Err(err).Msg("failed to update request")
		return err
	}

	return nil
}
