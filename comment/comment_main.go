package comment

import (
	_comHttpDelivery "github.com/P44elovod/task-management-app/comment/delivery/http"
	_comRepository "github.com/P44elovod/task-management-app/comment/repository/psql"
	_comUsecase "github.com/P44elovod/task-management-app/comment/usecase"
	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/project"
)

type CommentInit struct {
	CommentRepository domain.CommentRepository
	CommentUsecase    domain.CommentUseCase
}

func InitComment(initData *project.InitData) *CommentInit {
	cmr := _comRepository.NewPsqlCommentRepository(initData.DB)
	cmu := _comUsecase.NewCommentUsecase(cmr)

	_comHttpDelivery.New(initData.Router, initData.Logger, cmu, cmr)

	initData.Logger.Info("Comment Inited")

	return &CommentInit{
		CommentRepository: cmr,
		CommentUsecase:    cmu,
	}

}
