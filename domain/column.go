package domain

type Column struct {
	ID        uint
	ProjectID uint
	Name      string
	Position  uint
}

type ColumnUseCase interface {
	// CreateColumn()
	// UpdateColumnByID()
	// DeleteColumnByID()
	// FetchColumnsByProjectID()
}
type ColumnRepository interface{}
