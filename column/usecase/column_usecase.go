package columnusacase

import (
	"errors"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
)

type columnUsecase struct {
	columnRepo domain.ColumnRepository
	teaskRepo  domain.TaskRepository
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
		helpers.FailOnError(err, "Columnlist querying went wrong")
		return nil, err
	}

	for i := 0; i < len(columnList); i++ {
		taskList, err := c.teaskRepo.GetAllByColumnID(columnList[i].ID)
		if err != nil {
			helpers.FailOnError(err, "Tasklist querying went wrong")
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
		helpers.FailOnError(err, "Querying column went wrong")
		return err
	}

	isLast := c.columnRepo.CheckIfLastColumn(column.ProjectID)

	if column.Position != 1 {
		newColID, err = c.columnRepo.GetColumnIDByPositionAndProjectID(column.ProjectID, column.Position-1)
		if err != nil {
			helpers.FailOnError(err, "Querying columnID went wrong")
			return err
		}
	}

	if column.Position == 1 && isLast == false {
		newColID, err = c.columnRepo.GetColumnIDByPositionAndProjectID(column.ProjectID, column.Position+1)
		if err != nil {
			helpers.FailOnError(err, "Querying columnID went wrong")
			return err
		}
	}

	if column.Position == 1 && isLast == true {
		helpers.FailOnError(err, "The last column couldn't be deleted")
		return err
	}

	err = c.teaskRepo.UpdateColumnID(column.ID, newColID)
	if err != nil {
		helpers.FailOnError(err, "ColumnID updating went wrong")
		return err
	}

	err = c.columnRepo.DeleteByID(id)
	if err != nil {
		return err
	}

	return nil
}
