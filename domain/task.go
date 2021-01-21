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
	GetTaskWithCommentByID(id string) (Task, error)
	// DeleteTaskByID()
	// FetchTaskByID()
	// FetchTasksByColumnID()

}
type TaskRepository interface {
	StoreTask(task *Task) error
	GetByID(id string) (Task, error)
	GetAllByColumnID(id uint) ([]Task, error)
}
