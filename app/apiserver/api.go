package apiserver

import (
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

	initData := _project.InitData{
		Router: srv.router,
		Logger: srv.logger,
		DB:     db,
	}

	initEntities(&initData)

	srv.logger.Info("Server Started")
	if err := srv.start(); err != nil {
		srv.logger.Error(err)
		return err
	}
	return nil
}

func initEntities(initData *_project.InitData) {
	comment := _comment.InitComment(initData)
	task := _task.InitTask(initData, comment.CommentRepository)
	column := _column.InitColumn(initData, task.TaskRepository)

	_project.InitProject(initData, column.ColumnUsecase, task.TaskRepository)
}
