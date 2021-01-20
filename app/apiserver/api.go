package apiserver

import (
	"log"

	"github.com/P44elovod/task-management-app/config"
	"github.com/P44elovod/task-management-app/helpers"

	_column "github.com/P44elovod/task-management-app/column"
	_project "github.com/P44elovod/task-management-app/project"
)

type Api struct {
	server *Server
	db     *DB
}

func (a *Api) Start(config *config.Config) error {
	srv := a.server.newServer(config)
	db, err := a.db.newDB(config)
	helpers.FailOnError(err, "database connection doesn't work")

	_column.InitColumn(srv.router, db)
	_project.InitProject(srv.router, db)

	if err := srv.start(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
