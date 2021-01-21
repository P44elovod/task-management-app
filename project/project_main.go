package project

import (
	"database/sql"

	"github.com/P44elovod/task-management-app/domain"
	_pHttpDelivery "github.com/P44elovod/task-management-app/project/delivery/http"
	_pRepository "github.com/P44elovod/task-management-app/project/repository/psql"
	_pUsecase "github.com/P44elovod/task-management-app/project/usecase"
	"github.com/gorilla/mux"
)

func InitProject(r *mux.Router, db *sql.DB, cu domain.ColumnUsecase, tr domain.TaskRepository) {
	pr := _pRepository.NewPsqlProjectleRepository(db)
	pu := _pUsecase.NewProjectUsecase(pr, cu, tr)

	_pHttpDelivery.New(r, pu)

}
