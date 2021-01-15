package domain

type Comment struct {
	ID     uint
	TaskID uint
	Text   string
}

type CommentUseCase interface {
	// CreateComment()
	// UpdateCommentByID()
	// DeleteCommetByID()
	// FetchAllCometsByTaskID ()

}
type CommentRepository interface{}
