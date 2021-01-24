package domain

type Project struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Columns     []Column `json:"columns"`
}

type ProjectUsecase interface {
	Create(project *Project) error
	GetAll() ([]Project, error)
	GetByID(id string) (Project, error)
	DeleteByID(id string) error
}
type ProjectRepository interface {
	Store(project *Project) error
	GetAll() ([]Project, error)
	GetByID(id string) (Project, error)
	DeleteByID(id string) error
}
