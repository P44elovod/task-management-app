package project

import (
	"database/sql"

	"github.com/P44elovod/task-management-app/domain"
	_pHttpDelivery "github.com/P44elovod/task-management-app/project/delivery/http"
	_pRepository "github.com/P44elovod/task-management-app/project/repository/psql"
	_pUsecase "github.com/P44elovod/task-management-app/project/usecase"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type InitData struct {
	Router *mux.Router
	Logger *logrus.Logger
	DB     *sql.DB
}

func InitProject(initData *InitData, cu domain.ColumnUsecase, tr domain.TaskRepository) {
	pr := _pRepository.NewPsqlProjectleRepository(initData.DB)
	pu := _pUsecase.NewProjectUsecase(pr, cu, tr)

	_pHttpDelivery.New(initData.Router, initData.Logger, pu)

}
