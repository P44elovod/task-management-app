package domain

type Column struct {
	ID        uint
	ProjectID uint
	Name      string
	Position  uint
}

type ColumnUsecase interface {
	CreateColumn(column *Column) error
	// UpdateColumnByID()
	// DeleteColumnByID()
	// FetchColumnsByProjectID()
}
type ColumnRepository interface {
	StoreColumn(column *Column) error
}
