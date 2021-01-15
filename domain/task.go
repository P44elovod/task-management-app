package domain

type Task struct {
	ID          uint
	ColumnID    uint
	Priority    uint
	Name        string
	Description string
}

type TaskUseCase interface {
	// CreateTask()
	// UpdateTaskByID()
	// DeleteTaskByID()
	// FetchTaskByID()
	// FetchTasksByColumnID()

}
type TaskRepository interface{}
