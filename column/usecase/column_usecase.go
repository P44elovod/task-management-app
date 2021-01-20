package columnusacase

import (
	"errors"

	"github.com/P44elovod/task-management-app/domain"
)

type columnUsecase struct {
	columnRepo domain.ColumnRepository
}

func NewColumnUsecase(cr domain.ColumnRepository) domain.ColumnUsecase {
	return &columnUsecase{
		columnRepo: cr,
	}
}

func (c *columnUsecase) CreateColumn(column *domain.Column) error {

	if c.columnRepo.CheckColumnNameExists(&column.Name) == false {
		err := c.columnRepo.StoreColumn(column)
		if err != nil {
			return err
		}
	}

	return errors.New("Column name should be unique")
}

func (p *columnUsecase) FetchColumnsByProjectID(id string) ([]domain.Column, error) {

	projectsList, err := p.FetchColumnsByProjectID(id)
	if err != nil {
		return nil, err
	}

	return projectsList, nil
}
