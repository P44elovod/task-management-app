package taskusecase

import (
	"github.com/P44elovod/task-management-app/domain"
)

type taskUsecase struct {
	taskRepo    domain.TaskRepository
	commentRepo domain.CommentRepository
}

func NewTaskUsecase(tr domain.TaskRepository, cmr domain.CommentRepository) domain.TaskUseCase {
	return &taskUsecase{
		taskRepo:    tr,
		commentRepo: cmr,
	}
}

func (tu *taskUsecase) CreateTask(task *domain.Task) error {
	err := tu.taskRepo.StoreTask(task)
	if err != nil {
		return err
	}

	return nil
}

func (tu *taskUsecase) GetTaskWithCommentByID(id string) (domain.Task, error) {

	task, err := tu.taskRepo.GetByID(id)

	commentList, err := tu.commentRepo.GetAllByTaskID(id)
	if err != nil {
		return task, err
	}

	task.Comments = commentList

	return task, nil
}

func (tu *taskUsecase) DeleteByID(id string) error {
	err := tu.commentRepo.DeleteAllByTaskID(id)
	if err != nil {
		return err
	}

	err = tu.taskRepo.DeleteByID(id)
	if err != nil {
		return err
	}

	return nil
}

func (tu *taskUsecase) Update(task *domain.Task) error {
	err := tu.taskRepo.Update(task)
	if err != nil {
		return err
	}

	return nil
}
