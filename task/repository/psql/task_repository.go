package taskpsqlrepository

import (
	"database/sql"

	"github.com/P44elovod/task-management-app/domain"
)

type psqlTaskRepository struct {
	db *sql.DB
}

func NewPsqlTaskRepository(db *sql.DB) domain.TaskRepository {
	return &psqlTaskRepository{db: db}
}

func (tr *psqlTaskRepository) StoreTask(task *domain.Task) error {
	tx, err := tr.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO task (name, description, column_id, position) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	stmt.QueryRow(task.Name, task.Description, task.ColumnID, task.Priority).Scan(&task.ID)
	return tx.Commit()
}
