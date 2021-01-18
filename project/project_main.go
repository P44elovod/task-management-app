package project

import (
	"database/sql"

	_pHttpDelivery "github.com/P44elovod/task-management-app/project/delivery/http"
	_pRepository "github.com/P44elovod/task-management-app/project/repository/psql"
	_pUsecase "github.com/P44elovod/task-management-app/project/usecase"
	"github.com/gorilla/mux"
)

func InitProject(r *mux.Router, db *sql.DB) {
	pr := _pRepository.NewPsqlProjectleRepository(db)
	pu := _pUsecase.NewProjectUsecase(pr)

	_pHttpDelivery.New(r, pu)
}
