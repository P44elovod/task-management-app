package column

import (
	"database/sql"

	_cHttpDelivery "github.com/P44elovod/task-management-app/column/delivery/http"
	_cRepository "github.com/P44elovod/task-management-app/column/repository/psql"
	_cUsecase "github.com/P44elovod/task-management-app/column/usecase"
	"github.com/P44elovod/task-management-app/domain"
	"github.com/gorilla/mux"
)

type ColumnInit struct {
	ColumnRepository domain.ColumnRepository
	ColumnUsecase    domain.ColumnUsecase
}

func InitColumn(r *mux.Router, db *sql.DB) *ColumnInit {
	cr := _cRepository.NewPsqlColumnRepository(db)
	cu := _cUsecase.NewColumnUsecase(cr)

	_cHttpDelivery.New(r, cu)

	return &ColumnInit{
		ColumnRepository: cr,
		ColumnUsecase:    cu,
	}
}
