package domain

type Comment struct {
	ID     uint   `json:"id"`
	TaskID uint   `json:"task_id"`
	Text   string `json:"text"`
}

type CommentUseCase interface{}

type CommentRepository interface {
	StoreComment(comment *Comment) error
	GetAllByTaskID(id uint) ([]Comment, error)
	DeleteByID(id uint) error
	DeleteAllByTaskID(id uint) error
	UpdateByID(comment *Comment) error
}
