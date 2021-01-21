package comentpsqlrepository

import (
	"database/sql"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
)

type psqlCommentRepository struct {
	db *sql.DB
}

func NewPsqlCommentRepository(db *sql.DB) domain.CommentRepository {
	return &psqlCommentRepository{db: db}
}

func (cmr *psqlCommentRepository) StoreComment(comment *domain.Comment) error {
	tx, err := cmr.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO comment (task_id, text) VALUES ($1, $2) RETURNING id")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	stmt.QueryRow(comment.TaskID, comment.Text).Scan(&comment.ID)
	return tx.Commit()
}

func (cmr *psqlCommentRepository) GetAllByTaskID(id string) ([]domain.Comment, error) {

	rows, err := cmr.db.Query("SELECT id, text, task_id FROM comment WHERE task_id=$1 ORDER BY created_at", id)
	if err != nil {
		helpers.FailOnError(err, "Comment DB query processing went wrong!")
		return nil, err
	}
	var commentList []domain.Comment
	for rows.Next() {
		comment := domain.Comment{}
		err = rows.Scan(&comment.ID, &comment.Text, &comment.TaskID)
		if err != nil {
			helpers.FailOnError(err, " Comment DB row deserialization went wrong!")
			return nil, err
		}

		commentList = append(commentList, comment)
	}
	return commentList, nil
}