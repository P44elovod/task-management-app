package columnpsqlrepository

import (
	"database/sql"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
)

type psqlColumnRepository struct {
	db *sql.DB
}

func NewPsqlColumnRepository(db *sql.DB) domain.ColumnRepository {
	return &psqlColumnRepository{db: db}
}

func (cr *psqlColumnRepository) FetchColumnsByProjectID(id string) ([]domain.Column, error) {
	rows, err := cr.db.Query("SELECT id, name, position FROM column WHERE project_id=$1", id)
	if err != nil {
		helpers.FailOnError(err, "DB query processing went wrong!")
		return nil, err
	}
	var columnList []domain.Column
	for rows.Next() {
		column := domain.Column{}
		err = rows.Scan(&column.ID, &column.Name, &column.Position)
		if err != nil {
			helpers.FailOnError(err, "DB row deserialization went wrong!")
			return nil, err
		}

		columnList = append(columnList, column)
	}
	return columnList, nil
}

func (cr *psqlColumnRepository) StoreColumn(column *domain.Column) error {

	tx, err := cr.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO project_column (name, project_id, position) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	stmt.QueryRow(column.Name, column.ProjectID, column.Position).Scan(&column.ID)
	return tx.Commit()

}

func (cr *psqlColumnRepository) CheckColumnNameExists(name *string) bool {
	var count int
	cr.db.QueryRow("SELECT COUNT(name) FROM project_column WHERE name=$1", name).Scan(&count)
	if count > 0 {
		return true
	}

	return false
}
