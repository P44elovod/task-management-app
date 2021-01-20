package projectusecase

import (
	"fmt"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
)

type projectUsecase struct {
	projectRepo   domain.ProjectRepository
	columnUsecase domain.ColumnUsecase
}

func NewProjectUsecase(pr domain.ProjectRepository, cu domain.ColumnUsecase) domain.ProjectUsecase {
	return &projectUsecase{
		projectRepo:   pr,
		columnUsecase: cu,
	}
}

func (p *projectUsecase) CreateProject(project *domain.Project) error {

	err := p.projectRepo.StoreProject(project)
	if err != nil {
		return err
	}

	defultColumnName := helpers.GenerateColumnName(project.Name)
	defaultColumn := domain.Column{
		ProjectID: project.ID,
		Name:      defultColumnName,
		Position:  1,
	}

	err = p.columnUsecase.CreateColumn(&defaultColumn)
	helpers.FailOnError(err, "Column storing went wrong")

	project.Columns = append(project.Columns, defaultColumn)

	return nil
}

func (p *projectUsecase) Fetch() {
	fmt.Println("Project USecase Fetch")
	p.projectRepo.GetProjectByID()
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
