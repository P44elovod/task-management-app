package projectusecase

import (
	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
)

type projectUsecase struct {
	projectRepo   domain.ProjectRepository
	columnUsecase domain.ColumnUsecase
	taskRepo      domain.TaskRepository
}

func NewProjectUsecase(pr domain.ProjectRepository, cu domain.ColumnUsecase, tr domain.TaskRepository) domain.ProjectUsecase {
	return &projectUsecase{
		projectRepo:   pr,
		taskRepo:      tr,
		columnUsecase: cu,
	}
}

func (p *projectUsecase) Create(project *domain.Project) error {

	err := p.projectRepo.Store(project)
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
	if err != nil {
		return err
	}

	project.Columns = append(project.Columns, defaultColumn)

	return nil
}

func (p *projectUsecase) GetAll() ([]domain.Project, error) {

	projectsList, err := p.projectRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return projectsList, nil
}

func (p *projectUsecase) GetByID(id string) (domain.Project, error) {

	project, err := p.projectRepo.GetByID(id)
	if err != nil {
		return project, err
	}

	columnList, err := p.columnUsecase.GetColumnsWithTasksByProjectID(id)
	if err != nil {
		return project, err
	}

	project.Columns = columnList

	return project, nil
}

func (p *projectUsecase) DeleteByID(id string) error {

	return nil
}
