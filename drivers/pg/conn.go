package pg

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PGDriver struct {
	DBHost string
	DB     *sql.DB
}

func NewDriver(host string) (*PGDriver, error) {
	db, err := sql.Open("postgres", host)
	if err != nil {
		return nil, err
	}

	return &PGDriver{host, db}, nil
}

func (d *PGDriver) Host() string {
	return d.DBHost
}
