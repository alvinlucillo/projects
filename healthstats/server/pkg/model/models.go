package model

import "database/sql"

type Request struct {
	ID           string `db:"id"`
	FileName     string `db:"file_name"`
	Status       string `db:"status"`
	ErrorDetails string `db:"error_details"`
	Dates
}

type Dates struct {
	CreatedAt string         `db:"created_at"`
	UpdatedAt string         `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}
