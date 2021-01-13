package domain

type Project struct {
	ID   uint
	Name string
}

type ProjectUseCase interface {
	CreateProject() (uint, error)
	FetchAllProjects() ([]Project, error)
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
