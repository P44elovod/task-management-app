package domain

type Comment struct {
	ID     uint   `json:"id"`
	TaskID uint   `json:"task_id"`
	Text   string `json:"text"`
}

type CommentUseCase interface {
	CreateComment(comment *Comment) error
	// UpdateCommentByID()
	// DeleteCommetByID()
	// FetchAllCometsByTaskID ()

}
type CommentRepository interface {
	StoreComment(comment *Comment) error
	GetAllByTaskID(id string) ([]Comment, error)
}
