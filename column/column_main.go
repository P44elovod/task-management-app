package column

import (
	"database/sql"

	_cHttpDelivery "github.com/P44elovod/task-management-app/column/delivery/http"
	_cRepository "github.com/P44elovod/task-management-app/column/repository/psql"
	_cUsecase "github.com/P44elovod/task-management-app/column/usecase"
	"github.com/gorilla/mux"
)

func InitColumn(r *mux.Router, db *sql.DB) {
	cr := _cRepository.NewPsqlColumnRepository(db)
	cu := _cUsecase.NewColumnUsecase(cr)

	_cHttpDelivery.New(r, cu)
}
