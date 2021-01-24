package task

import (
	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/project"
	_tHttpDelivery "github.com/P44elovod/task-management-app/task/delivery/http"
	_tRepository "github.com/P44elovod/task-management-app/task/repository/psql"
	_tUsecase "github.com/P44elovod/task-management-app/task/usecase"
)

type TaskInit struct {
	TaskRepository domain.TaskRepository
	TaskUsecase    domain.TaskUseCase
}

func InitTask(initData *project.InitData, cmr domain.CommentRepository) *TaskInit {
	tr := _tRepository.NewPsqlTaskRepository(initData.DB)
	tu := _tUsecase.NewTaskUsecase(tr, cmr)

	_tHttpDelivery.New(initData.Router, initData.Logger, tu)

	return &TaskInit{
		TaskRepository: tr,
		TaskUsecase:    tu,
	}

}
