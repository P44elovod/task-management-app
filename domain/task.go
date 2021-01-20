package domain

type Task struct {
	ID          uint      `json:"id"`
	ColumnID    uint      `json:"column_id"`
	Priority    uint      `json:"priority"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Comments    []Comment `json:"comments"`
}

type TaskUseCase interface {
	CreateTask(task *Task) error
	// UpdateTaskByID()
	// DeleteTaskByID()
	// FetchTaskByID()
	// FetchTasksByColumnID()

}
type TaskRepository interface {
	StoreTask(task *Task) error
}
