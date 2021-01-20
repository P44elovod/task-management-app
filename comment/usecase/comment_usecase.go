package commentusecase

import "github.com/P44elovod/task-management-app/domain"

type commentUsecase struct {
	commentRepo domain.CommentRepository
}

func NewCommentUsecase(cmr domain.CommentRepository) domain.CommentUseCase {
	return &commentUsecase{
		commentRepo: cmr,
	}
}

func (cmu *commentUsecase) CreateComment(comment *domain.Comment) error {
	err := cmu.commentRepo.StoreComment(comment)
	if err != nil {
		return err
	}
	return nil
}
