package domain

type Comment struct {
	ID     uint
	TaskID uint
	Body   string
}

type CommentUseCase interface{}
type CommentRepository interface{}
