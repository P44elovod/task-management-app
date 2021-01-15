package projectusecase

import (
	"encoding/json"
	"fmt"
	"sort"

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

func (p *projectUsecase) FetchAllProjects() ([]byte, error) {

	projectsList, err := p.projectRepo.FetchAllProjects()
	if err != nil {
		return nil, err
	}
	sort.Slice(projectsList, func(i, j int) bool {
		return projectsList[i].Name < projectsList[j].Name
	})

	jList, err := json.Marshal(projectsList)

	return jList, nil
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
