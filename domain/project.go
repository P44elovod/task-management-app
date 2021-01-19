package domain

type Project struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectUseCase interface {
	CreateProject(project *Project) error
	FetchAllProjects() ([]Project, error)
	GetProjectByID() (Project, error)
	DeleteProject() error
	DeleteAllProjects() error
	UpdateProject() (uint, error)
}
type ProjectRepository interface {
	FetchAllProjects() ([]Project, error)
	GetProjectByID() (Project, error)
	StoreProject(project *Project) error
	UpdateProject() (uint, error)
	DeleteAllProjects() error
	DeleteProject() error
}
