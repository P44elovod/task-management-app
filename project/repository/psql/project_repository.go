package projectpsqlrepository

import (
	"database/sql"
	"fmt"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
)

type psqlProjectRepository struct {
	db *sql.DB
}

func NewPsqlProjectleRepository(db *sql.DB) domain.ProjectRepository {
	return &psqlProjectRepository{db}
}

func (p *psqlProjectRepository) FetchAllProjects() ([]domain.Project, error) {
	rows, err := p.db.Query("SELECT id, name, description FROM project ORDER BY name")
	if err != nil {
		helpers.FailOnError(err, "DB query processing went wrong!")
		return nil, err
	}
	var progectList []domain.Project
	for rows.Next() {
		project := domain.Project{}
		err = rows.Scan(&project.ID, &project.Name, &project.Description)
		if err != nil {
			helpers.FailOnError(err, "DB row deserialization went wrong!")
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

func (p *psqlProjectRepository) UpdateProject() (uint, error) {
	return 0, nil
}
func (p *psqlProjectRepository) DeleteAllProjects() error {
	return nil
}
func (p *psqlProjectRepository) DeleteProject() error {
	return nil
}

func (p *psqlProjectRepository) GetProjectByID() (domain.Project, error) {

	fmt.Println("psqlProjectRepository fetch")
	return domain.Project{}, nil
}
