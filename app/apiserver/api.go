package apiserver

import (
	"database/sql"
	"log"

	"github.com/P44elovod/task-management-app/config"
	_project "github.com/P44elovod/task-management-app/project"
)

type Api struct {
	server Server
	DB     *sql.DB
}

func (a *Api) Start(config *config.Config) error {
	srv := a.server.newServer(config)

	_project.InitProject(srv.router, a.DB)

	if err := srv.start(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
