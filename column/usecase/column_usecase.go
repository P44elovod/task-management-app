package columnusacase

import "github.com/P44elovod/task-management-app/domain"

type columnUsecase struct {
	columnRepo domain.ColumnRepository
}

func NewColumnUsecase(cr domain.ColumnRepository) domain.ColumnUsecase {
	return &columnUsecase{
		columnRepo: cr,
	}
}

func (c *columnUsecase) CreateColumn(column *domain.Column) error {
	err := c.columnRepo.StoreColumn(column)
	if err != nil {
		return err
	}

	return nil
}
