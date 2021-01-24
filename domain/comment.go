package domain

type Comment struct {
	ID     uint   `json:"id"`
	TaskID uint   `json:"task_id"`
	Text   string `json:"text"`
}

type CommentUseCase interface{}

type CommentRepository interface {
	StoreComment(comment *Comment) error
	GetAllByTaskID(id string) ([]Comment, error)
	DeleteByID(id string) error
	DeleteAllByTaskID(id string) error
	UpdateByID(comment *Comment) error
}
