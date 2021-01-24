package projectpsqlrepository

import (
	"database/sql"

	"github.com/P44elovod/task-management-app/domain"
)

type psqlProjectRepository struct {
	db *sql.DB
}

func NewPsqlProjectleRepository(db *sql.DB) domain.ProjectRepository {
	return &psqlProjectRepository{
		db: db,
	}
}

func (p *psqlProjectRepository) FetchAllProjects() ([]domain.Project, error) {
	rows, err := p.db.Query("SELECT id, name, description FROM project ORDER BY name")
	if err != nil {
		return nil, err
	}
	var progectList []domain.Project
	for rows.Next() {
		project := domain.Project{}
		err = rows.Scan(&project.ID, &project.Name, &project.Description)
		if err != nil {
			return nil, err
		}

		progectList = append(progectList, project)
	}
	return progectList, nil
}

func (p *psqlProjectRepository) StoreProject(project *domain.Project) error {

	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO project (name, description) VALUES ($1, $2) RETURNING id")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	stmt.QueryRow(project.Name, project.Description).Scan(&project.ID)
	return tx.Commit()
}

func (p *psqlProjectRepository) GetProjectByID(id string) (domain.Project, error) {

	var project domain.Project
	row, err := p.db.Query("SELECT id, name, description FROM project WHERE id = $1", id)
	if err != nil {
		return project, err
	}

	for row.Next() {
		err = row.Scan(&project.ID, &project.Name, &project.Description)
		if err != nil {
			return project, err
		}

	}
	return project, nil
}

func (p *psqlProjectRepository) DeleteProjectByID(id string) error {

	_, err := p.db.Exec("DELETE FROM project WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (p *psqlProjectRepository) UpdateByID(project *domain.Project) error {

	_, err := p.db.Exec("UPDATE task SET name=$1, description=$2,  WHERE id=$3", project.Name, project.Description, project.ID)

	if err != nil {
		return err
	}

	return nil
}
