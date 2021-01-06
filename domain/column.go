package domain

type Column struct {
	ID        uint
	ProjectID uint
	Position  uint
	Name      string
}

type ColumnUseCase interface{}
type ColumnRepository interface{}
