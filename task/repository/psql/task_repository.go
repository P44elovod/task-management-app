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

func (tr *psqlTaskRepository) GetByID(id uint) (domain.Task, error) {

	var task domain.Task

	rows, err := tr.db.Query("SELECT id, name, description, column_id, position FROM task WHERE id = $1", id)
	if err != nil {
		return task, err
	}

	for rows.Next() {
		err = rows.Scan(&task.ID, &task.Description, &task.Name, &task.ColumnID, &task.Priority)
		if err != nil {
			return task, err
		}
	}

	return task, nil
}

func (tr *psqlTaskRepository) GetAllByColumnID(id uint) ([]domain.Task, error) {
	var taskList []domain.Task

	row, err := tr.db.Query("SELECT id, name, description, column_id, position, created_at FROM task WHERE column_id=$1 ORDER BY position", id)
	if err != nil {
		return taskList, err
	}

	for row.Next() {
		task := domain.Task{}
		err = row.Scan(&task.ID, &task.Name, &task.Description, &task.ColumnID, &task.Priority, &task.CreatedAt)
		if err != nil {
			return nil, err
		}

		taskList = append(taskList, task)
	}
	return taskList, nil
}

func (tr *psqlTaskRepository) DeleteByID(id uint) error {

	_, err := tr.db.Exec("DELETE FROM task WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (tr *psqlTaskRepository) DeleteByColumnID(id uint) error {

	_, err := tr.db.Exec("DELETE FROM task WHERE column_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (tr *psqlTaskRepository) Update(task *domain.Task) error {

	_, err := tr.db.Exec("UPDATE task SET name=$1, description=$2,  column_id=$3, position=$4 WHERE id=$5",
		task.Name,
		task.Description,
		task.ColumnID,
		task.Priority,
		task.ID)

	if err != nil {
		return err
	}

	return nil
}

func (tr *psqlTaskRepository) UpdateColumnID(oldColID, newColID uint) error {
	_, err := tr.db.Exec("UPDATE task SET column_id=$1 WHERE column_id=$2", newColID, oldColID)

	if err != nil {
		return err
	}

	return nil
}

func (tr *psqlTaskRepository) UpdatePriority(id, priority uint) error {
	_, err := tr.db.Exec("UPDATE task SET position=$1 WHERE id=$2", priority, id)

	return err
}
