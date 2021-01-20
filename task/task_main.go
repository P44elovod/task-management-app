package task

import (
	"database/sql"

	_tHttpDelivery "github.com/P44elovod/task-management-app/task/delivery/http"
	_tRepository "github.com/P44elovod/task-management-app/task/repository/psql"
	_tUsecase "github.com/P44elovod/task-management-app/task/usecase"
	"github.com/gorilla/mux"
)

func InitTask(r *mux.Router, db *sql.DB) {
	tr := _tRepository.NewPsqlTaskRepository(db)
	tu := _tUsecase.NewTaskUsecase(tr)

	_tHttpDelivery.New(r, tu)

}
