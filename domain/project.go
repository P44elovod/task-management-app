package domain

type Project struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectUseCase interface {
	CreateProject() (uint, error)
	FetchAllProjects() ([]byte, error)
	FetchProjectByID() (Project, error)
	DeleteProject() error
	DeleteAllProjects() error
	UpdateProject() (uint, error)
}
type ProjectRepository interface {
	FetchAllProjects() ([]Project, error)
	FetchProjectByID() (Project, error)
	StoreProject() (uint, error)
	UpdateProject() (uint, error)
	DeleteAllProjects() error
	DeleteProject() error
}
