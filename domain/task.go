package domain

import "time"

type Task struct {
	ID          uint      `json:"id"`
	ColumnID    uint      `json:"column_id"`
	Priority    uint      `json:"priority"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Comments    []Comment `json:"comments"`
}

type TaskUseCase interface {
	CreateTask(task *Task) error
	GetTaskWithCommentByID(id uint) (Task, error)
	DeleteByID(id uint) error
	Update(taskList *Task) error
	UpdateUpdatePriority(positionsList map[uint]uint) error
	PrepareShiftDownPriorityMap(columnID, startPosition uint) (map[uint]uint, error)
	PrepareShiftUpPriorityMap(columnID, startPosition uint) (map[uint]uint, error)
}
type TaskRepository interface {
	StoreTask(task *Task) error
	GetByID(id uint) (Task, error)
	GetAllByColumnID(id uint) ([]Task, error)
	DeleteByID(id uint) error
	DeleteByColumnID(id uint) error
	Update(taskList *Task) error
	UpdateColumnID(oldColID, newColID uint) error
	UpdatePriority(id, priority uint) error
}
