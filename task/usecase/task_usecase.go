package taskusecase

import "github.com/P44elovod/task-management-app/domain"

type taskUsecase struct {
	taskRepo domain.TaskRepository
}

func NewTaskUsecase(tr domain.TaskRepository) domain.TaskUseCase {
	return &taskUsecase{
		taskRepo: tr,
	}
}

func (tu *taskUsecase) CreateTask(task *domain.Task) error {
	err := tu.taskRepo.StoreTask(task)
	if err != nil {
		return err
	}

	return nil
}
