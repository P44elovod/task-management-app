package projectpsqlrepository

import (
	"database/sql"
	"fmt"

	"github.com/P44elovod/task-management-app/domain"
)

type psqlProjectRepository struct {
	Conn *sql.DB
}

func NewPsqlProjectleRepository(Conn *sql.DB) domain.ProjectRepository {
	return &psqlProjectRepository{Conn}
}

func (p *psqlProjectRepository) FetchAllProjects() ([]domain.Project, error) {
	return nil, nil
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
