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
	rows, err := p.db.Query("SELECT id, name, description FROM project")
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
func (p *psqlProjectRepository) FetchProjectByID() (domain.Project, error) {

	fmt.Println("psqlProjectRepository fetch")
	return domain.Project{}, nil
}
func (p *psqlProjectRepository) StoreProject() (uint, error) {
	return 0, nil
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
