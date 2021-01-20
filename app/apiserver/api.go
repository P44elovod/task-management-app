package apiserver

import (
	"log"

	"github.com/P44elovod/task-management-app/config"
	"github.com/P44elovod/task-management-app/helpers"

	_column "github.com/P44elovod/task-management-app/column"
	_comment "github.com/P44elovod/task-management-app/comment"
	_project "github.com/P44elovod/task-management-app/project"
	_task "github.com/P44elovod/task-management-app/task"
)

type Api struct {
	server *Server
	db     *DB
}

func (a *Api) Start(config *config.Config) error {
	srv := a.server.newServer(config)
	db, err := a.db.newDB(config)
	helpers.FailOnError(err, "database connection doesn't work")

	column := _column.InitColumn(srv.router, db)

	_project.InitProject(srv.router, db, column.ColumnUsecase)
	_task.InitTask(srv.router, db)
	_comment.InitComment(srv.router, db)

	if err := srv.start(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
