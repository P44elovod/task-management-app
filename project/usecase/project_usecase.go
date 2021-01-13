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
	p.projectRepo.Fetch()
}
