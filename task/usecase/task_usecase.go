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
	tasks, err := tu.taskRepo.GetAllByColumnID(task.ColumnID)
	if err != nil {
		return err
	}

	if len(tasks) >= 1 && task.Priority <= tasks[len(tasks)-1].Priority {

		priority, err := tu.PrepareShiftDownPriorityMap(task.ColumnID, task.Priority)
		if err != nil {
			return err
		}

		err = tu.UpdateUpdatePriority(priority)
		if err != nil {
			return err
		}

	}

	err = tu.taskRepo.StoreTask(task)
	if err != nil {
		return err
	}

	return nil
}

func (tu *taskUsecase) GetTaskWithCommentByID(id uint) (domain.Task, error) {

	task, err := tu.taskRepo.GetByID(id)

	commentList, err := tu.commentRepo.GetAllByTaskID(id)
	if err != nil {
		return task, err
	}

	task.Comments = commentList

	return task, nil
}

func (tu *taskUsecase) DeleteByID(id uint) error {
	// err := tu.commentRepo.DeleteAllByTaskID(id)
	// if err != nil {
	// 	return err
	// }
	task, err := tu.taskRepo.GetByID(id)
	if err != nil {
		return err
	}

	newPriority, err := tu.PrepareShiftUpPriorityMap(task.ColumnID, task.Priority)
	err = tu.UpdateUpdatePriority(newPriority)
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

func (tu *taskUsecase) PrepareShiftDownPriorityMap(columnID, startPosition uint) (map[uint]uint, error) {
	priorityMap := make(map[uint]uint)

	tasks, err := tu.taskRepo.GetAllByColumnID(columnID)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].Priority >= startPosition {
			priorityMap[tasks[i].ID] = tasks[i].Priority + 1
		}

	}

	return priorityMap, nil
}

func (tu *taskUsecase) PrepareShiftUpPriorityMap(columnID, startPosition uint) (map[uint]uint, error) {
	priorityMap := make(map[uint]uint)

	tasks, err := tu.taskRepo.GetAllByColumnID(columnID)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].Priority >= startPosition {
			priorityMap[tasks[i].ID] = tasks[i].Priority - 1
		}

	}

	return priorityMap, nil
}

func (tu *taskUsecase) UpdateUpdatePriority(positionsList map[uint]uint) error {

	for id, priority := range positionsList {
		err := tu.taskRepo.UpdatePriority(id, priority)
		if err != nil {
			return err
		}
	}

	return nil
}
