package domain

type Project struct {
	ID   uint
	Name string
}

type ProjectUseCase interface {
	Fetch()
}
type ProjectRepository interface {
	Fetch()
}
