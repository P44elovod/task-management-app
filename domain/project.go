package domain

type Project struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Columns     []Column `json:"columns"`
}

type ProjectUsecase interface {
	CreateProject(project *Project) error
	FetchAllProjects() ([]Project, error)
	GetProjectByID(id string) (Project, error)
}
type ProjectRepository interface {
	FetchAllProjects() ([]Project, error)
	GetProjectByID(id string) (Project, error)
	StoreProject(project *Project) error
	DeleteProjectByID(id string) error
}
