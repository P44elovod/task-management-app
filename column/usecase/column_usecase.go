package columnusacase

import (
	"errors"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/sirupsen/logrus"
)

type columnUsecase struct {
	columnRepo domain.ColumnRepository
	teaskRepo  domain.TaskRepository
	logger     *logrus.Logger
}

func NewColumnUsecase(cr domain.ColumnRepository, tr domain.TaskRepository) domain.ColumnUsecase {
	return &columnUsecase{
		columnRepo: cr,
		teaskRepo:  tr,
	}
}

func (c *columnUsecase) CreateColumn(column *domain.Column) error {

	if c.columnRepo.CheckColumnNameExists(&column.Name) == false {
		err := c.columnRepo.StoreColumn(column)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("Column name should be unique")
}

func (c *columnUsecase) GetColumnsWithTasksByProjectID(id string) ([]domain.Column, error) {

	columnList, err := c.columnRepo.GetColumnsByProjectID(id)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(columnList); i++ {
		taskList, err := c.teaskRepo.GetAllByColumnID(columnList[i].ID)
		if err != nil {
			return nil, err
		}
		columnList[i].Tasks = taskList
	}

	return columnList, nil
}

func (c *columnUsecase) DeleteByID(id string) error {

	var newColID uint

	column, err := c.columnRepo.GetByID(id)
	if err != nil {
		return err
	}

	isLast := c.columnRepo.CheckIfLastColumn(column.ProjectID)

	if column.Position != 1 {
		newColID, err = c.columnRepo.GetColumnIDByPositionAndProjectID(column.ProjectID, column.Position-1)
		if err != nil {
			return err
		}
	}

	if column.Position == 1 && isLast == false {
		newColID, err = c.columnRepo.GetColumnIDByPositionAndProjectID(column.ProjectID, column.Position+1)
		if err != nil {
			return err
		}
	}

	if column.Position == 1 && isLast == true {
		return err
	}

	err = c.teaskRepo.UpdateColumnID(column.ID, newColID)
	if err != nil {
		return err
	}

	err = c.columnRepo.DeleteByID(id)
	if err != nil {
		return err
	}

	return nil
}
