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
	GetColumnsWithTasksByProjectID(id uint) ([]Column, error)
	DeleteByID(id uint) error
	Update(column *Column) error
	UpdatePosition(positionsList map[uint]uint) error
}
type ColumnRepository interface {
	StoreColumn(column *Column) error
	CheckColumnNameExists(name *string) bool
	GetColumnsByProjectID(id uint) ([]Column, error)
	GetByID(id uint) (Column, error)
	GetColumnIDByPositionAndProjectID(id, position uint) (uint, error)
	DeleteByID(id uint) error
	CheckIfLastColumn(projectID uint) bool
	Update(column *Column) error
	UpdatePositions(id, position uint) error
}
