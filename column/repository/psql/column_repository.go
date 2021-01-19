package columnpsqlrepository

import (
	"database/sql"
	"errors"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
)

type psqlColumnRepository struct {
	db *sql.DB
}

func NewPsqlColumnRepository(db *sql.DB) domain.ColumnRepository {
	return &psqlColumnRepository{db}
}

func (p *psqlColumnRepository) StoreColumn(column *domain.Column) error {

	if p.CheckColumnNameExists(column.Name) {
		return errors.New("Column name should be unique")
	}
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO column (name, project_id) VALUES ($1, $2) RETURNING id")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	stmt.QueryRow(column.Name, column.ProjectID).Scan(&column.ID)

	return tx.Commit()
}

func (p *psqlColumnRepository) CheckColumnNameExists(name string) bool {
	row := p.db.QueryRow("SELECT COUNT(name) FROM column WHERE name=$1", name)

	var count int
	err := row.Scan(&count)
	if err != nil {
		helpers.FailOnError(err, "Row Scan went wrong")
	}

	if count > 0 {
		return true
	}

	return false
}
