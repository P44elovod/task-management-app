package taskpsqlrepository

import (
	"database/sql"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
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

func (tr *psqlTaskRepository) GetByID(id string) (domain.Task, error) {

	var task domain.Task

	rows, err := tr.db.Query("SELECT id, name, description, column_id, position FROM task WHERE id = $1", id)
	if err != nil {
		helpers.FailOnError(err, "Task DB query processing went wrong!")
		return task, err
	}

	for rows.Next() {
		err = rows.Scan(&task.ID, &task.Description, &task.Name, &task.ColumnID, &task.Priority)
		if err != nil {
			helpers.FailOnError(err, "Task DB row deserialization went wrong!")
			return task, err
		}
	}

	return task, nil
}

func (tr *psqlTaskRepository) GetAllByColumnID(id uint) ([]domain.Task, error) {
	var taskList []domain.Task

	row, err := tr.db.Query("SELECT id, name, description, column_id, position, created_at FROM task WHERE column_id=$1 ORDER BY position", id)
	if err != nil {
		helpers.FailOnError(err, "Task DB query processing went wrong!")
		return taskList, err
	}

	for row.Next() {
		task := domain.Task{}
		err = row.Scan(&task.ID, &task.Name, &task.Description, &task.ColumnID, &task.Priority, &task.CreatedAt)
		if err != nil {
			helpers.FailOnError(err, "Task DB row deserialization went wrong!")
			return nil, err
		}

		taskList = append(taskList, task)
	}
	return taskList, nil
}
