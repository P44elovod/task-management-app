package domain

type Task struct {
	ID          uint
	ColumnID    uint
	Priority    uint
	Name        string
	Description string
}

type TaskUseCase interface{}
type TaskRepository interface{}
