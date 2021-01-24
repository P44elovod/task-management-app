package column

import (
	_cHttpDelivery "github.com/P44elovod/task-management-app/column/delivery/http"
	_cRepository "github.com/P44elovod/task-management-app/column/repository/psql"
	_cUsecase "github.com/P44elovod/task-management-app/column/usecase"
	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/project"
)

type ColumnInit struct {
	ColumnRepository domain.ColumnRepository
	ColumnUsecase    domain.ColumnUsecase
}

func InitColumn(initData *project.InitData, tr domain.TaskRepository) *ColumnInit {
	cr := _cRepository.NewPsqlColumnRepository(initData.DB)
	cu := _cUsecase.NewColumnUsecase(cr, tr)

	_cHttpDelivery.New(initData.Router, initData.Logger, cu)

	return &ColumnInit{
		ColumnRepository: cr,
		ColumnUsecase:    cu,
	}
}
