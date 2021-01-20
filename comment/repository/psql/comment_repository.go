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
