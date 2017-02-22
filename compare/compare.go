package compare

import "database/sql"

type (
	Driver interface {
		Count() error
		Views() error
	}

	Comparison func(*sql.DB, *sql.DB) error
)
