package comentpsqlrepository

import (
	"database/sql"

	"github.com/P44elovod/task-management-app/domain"
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

	rows, err := cmr.db.Query("SELECT id, text, task_id FROM comment WHERE task_id=$1 ORDER BY created_at DESC", id)
	if err != nil {
		return nil, err
	}
	var commentList []domain.Comment
	for rows.Next() {
		comment := domain.Comment{}
		err = rows.Scan(&comment.ID, &comment.Text, &comment.TaskID)
		if err != nil {
			return nil, err
		}

		commentList = append(commentList, comment)
	}
	return commentList, nil
}

func (cmr *psqlCommentRepository) DeleteByID(id string) error {
	_, err := cmr.db.Exec("DELETE FROM comment WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (cmr *psqlCommentRepository) DeleteAllByTaskID(id string) error {

	_, err := cmr.db.Exec("DELETE FROM comment WHERE task_id=$1 RETURNING id", id)
	if err != nil {
		return err
	}
	return nil
}

func (cmr *psqlCommentRepository) UpdateByID(comment *domain.Comment) error {
	_, err := cmr.db.Exec("UPDATE comment SET text=$1 WHERE id=$2", comment.Text, comment.ID)

	return err
}
