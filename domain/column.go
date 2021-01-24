package domain

type Column struct {
	ID        uint   `json:"id"`
	ProjectID uint   `json:"project_id"`
	Name      string `json:"name"`
	Position  uint   `json:"position"`
	Tasks     []Task `json:"tasks"`
}

type ColumnUsecase interface {
	CreateColumn(column *Column) error
	GetColumnsWithTasksByProjectID(id string) ([]Column, error)
	DeleteByID(id string) error
}
type ColumnRepository interface {
	StoreColumn(column *Column) error
	CheckColumnNameExists(name *string) bool
	GetColumnsByProjectID(id string) ([]Column, error)
	GetByID(id string) (Column, error)
	GetColumnIDByPositionAndProjectID(id, position uint) (uint, error)
	DeleteByID(id string) error
	CheckIfLastColumn(projectID uint) bool
}
