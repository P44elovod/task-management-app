package apiserver

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/P44elovod/task-management-app/config"
	"github.com/P44elovod/task-management-app/helpers"
)

type DB struct {
	config *config.Config
}

func (d *DB) newDB(config *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DBUrl)
	helpers.FailOnError(err, "Open database connection went wrong")
	pingError := db.Ping()

	if pingError != nil {
		return nil, pingError
	}

	return db, nil
}
