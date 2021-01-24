package comment

import (
	"database/sql"

	_comHttpDelivery "github.com/P44elovod/task-management-app/comment/delivery/http"
	_comRepository "github.com/P44elovod/task-management-app/comment/repository/psql"
	_comUsecase "github.com/P44elovod/task-management-app/comment/usecase"
	"github.com/P44elovod/task-management-app/domain"
	"github.com/gorilla/mux"
)

type CommentInit struct {
	CommentRepository domain.CommentRepository
	CommentUsecase    domain.CommentUseCase
}

func InitComment(r *mux.Router, db *sql.DB) *CommentInit {
	cmr := _comRepository.NewPsqlCommentRepository(db)
	cmu := _comUsecase.NewCommentUsecase(cmr)

	_comHttpDelivery.New(r, cmu, cmr)

	return &CommentInit{
		CommentRepository: cmr,
		CommentUsecase:    cmu,
	}

}
