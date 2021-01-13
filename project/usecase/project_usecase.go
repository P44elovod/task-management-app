package projectusecase

import (
	"fmt"

	"github.com/P44elovod/task-management-app/domain"
)

type projectUsecase struct {
	projectRepo domain.ProjectRepository
}

func NewArticleUsecase(pr domain.ProjectRepository) domain.ProjectUseCase {
	return &projectUsecase{
		projectRepo: pr,
	}
}

func (p *projectUsecase) Fetch() {
	fmt.Println("Project USecase Fetch")
	p.projectRepo.FetchProjectByID()
}

func (p *projectUsecase) CreateProject() (uint, error) {
	return 0, nil
}

func (p *projectUsecase) FetchAllProjects() ([]domain.Project, error) {
	return nil, nil
}

func (p *projectUsecase) FetchProjectByID() (domain.Project, error) {
	return domain.Project{}, nil
}

func (p *projectUsecase) DeleteProject() error {
	return nil
}

func (p *projectUsecase) DeleteAllProjects() error {
	return nil
}

func (p *projectUsecase) UpdateProject() (uint, error) {
	return 0, nil
}
