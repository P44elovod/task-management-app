package columnpsqlrepository

import (
	"database/sql"

	"github.com/P44elovod/task-management-app/domain"
)

type psqlColumnRepository struct {
	db *sql.DB
}

func NewPsqlColumnRepository(db *sql.DB) domain.ColumnRepository {
	return &psqlColumnRepository{db: db}
}

func (cr *psqlColumnRepository) GetColumnsByProjectID(id uint) ([]domain.Column, error) {
	rows, err := cr.db.Query("SELECT id, name, position, project_id FROM project_column WHERE project_id=$1 ORDER BY position", id)
	if err != nil {
		return nil, err
	}
	var columnList []domain.Column
	for rows.Next() {
		column := domain.Column{}
		err = rows.Scan(&column.ID, &column.Name, &column.Position, &column.ProjectID)
		if err != nil {
			return nil, err
		}

		columnList = append(columnList, column)
	}
	return columnList, nil
}

func (cr *psqlColumnRepository) GetByID(id uint) (domain.Column, error) {

	var column domain.Column

	rows, err := cr.db.Query("SELECT id, name, position, project_id FROM project_column WHERE id=$1", id)
	if err != nil {
		return column, err
	}

	for rows.Next() {
		err = rows.Scan(&column.ID, &column.Name, &column.Position, &column.ProjectID)
		if err != nil {
			return column, err
		}

	}
	return column, nil
}

func (cr *psqlColumnRepository) GetColumnIDByPositionAndProjectID(id, position uint) (uint, error) {

	var ID uint

	rows, err := cr.db.Query("SELECT id FROM project_column WHERE project_id=$1 AND position=$2", id, position)
	if err != nil {
		return ID, err
	}

	for rows.Next() {
		err = rows.Scan(&ID)
		if err != nil {
			return ID, err
		}

	}
	return ID, nil
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

func (cr *psqlColumnRepository) DeleteByID(id uint) error {
	_, err := cr.db.Exec("DELETE FROM project_column WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (cr *psqlColumnRepository) CheckColumnNameExists(name *string) bool {
	var count int
	cr.db.QueryRow("SELECT COUNT(name) FROM project_column WHERE name=$1", name).Scan(&count)
	if count > 0 {
		return true
	}

	return false
}

func (cr *psqlColumnRepository) CheckIfLastColumn(projectID uint) bool {

	var count int
	cr.db.QueryRow("Select COUNT(id) FROM project_column WHERE project_id=$1", projectID).Scan(&count)

	if count > 1 {
		return false
	}

	return true
}

func (cr *psqlColumnRepository) Update(column *domain.Column) error {
	_, err := cr.db.Exec("UPDATE project_column SET name=$1, position=$2 WHERE id=$3", column.Name, column.Position, column.ID)

	return err
}

func (cr *psqlColumnRepository) UpdatePositions(id, position uint) error {
	_, err := cr.db.Exec("UPDATE project_column SET position=$1 WHERE id=$2", position, id)

	return err
}
