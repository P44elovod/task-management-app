package domain

type Column struct {
	ID        uint   `json:"id"`
	ProjectID uint   `json:"project_id"`
	Name      string `json:"name"`
	Position  uint   `json:"position"`
}

type ColumnUsecase interface {
	CreateColumn(column *Column) error
	FetchColumnsByProjectID(id string) ([]Column, error)
	// UpdateColumnByID()
	// DeleteColumnByID()
}
type ColumnRepository interface {
	StoreColumn(column *Column) error
	CheckColumnNameExists(name *string) bool
	FetchColumnsByProjectID(id string) ([]Column, error)
}
