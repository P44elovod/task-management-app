package projectusecase

import (
	"fmt"

	"github.com/P44elovod/task-management-app/domain"
)

type projectUsecase struct {
	projectRepo domain.ProjectRepository
}

func NewProjectUsecase(pr domain.ProjectRepository) domain.ProjectUseCase {
	return &projectUsecase{
		projectRepo: pr,
	}
}

func (p *projectUsecase) Fetch() {
	fmt.Println("Project USecase Fetch")
	p.projectRepo.GetProjectByID()
}

func (p *projectUsecase) CreateProject(project *domain.Project) error {
	err := p.projectRepo.StoreProject(project)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectUsecase) FetchAllProjects() ([]domain.Project, error) {

	projectsList, err := p.projectRepo.FetchAllProjects()
	if err != nil {
		return nil, err
	}

	return projectsList, nil
}

func (p *projectUsecase) GetProjectByID() (domain.Project, error) {
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
