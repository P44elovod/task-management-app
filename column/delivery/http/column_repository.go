package columnpsqlrepository

import (
	"database/sql"

	"github.com/P44elovod/task-management-app/domain"
)

type psqlColumnRepository struct {
	db *sql.DB
}

func NewPsqlColumnRepository(db *sql.DB) domain.ColumnRepository {
	return &psqlColumnRepository{db}
}

func (p *psqlColumnRepository) StoreColumn(column *domain.Column) error {
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
