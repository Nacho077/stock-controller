package repository

import (
	"fmt"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/types"
)

type MovementRepositoryInterface interface {
	GetMovementsByCompany(id int) ([]types.Movement, error)
}

func (repository Repository) GetMovementsByCompany(id int) ([]types.Movement, error) {
	var movements []types.Movement

	company, err := repository.Db.Query("SELECT * FROM company WHERE company.id = ?", id)
	if err != nil {
		errors.NewFailedDependencyError(fmt.Sprintf("Error in database when bringing company with id %d", id), err.Error())
	}

	if company == nil {
		errors.NewBadRequestError(fmt.Sprintf("Error when searching for a company, the company with id %d doesn't exist", id), "User error")
	}

	fmt.Println("LA RECONCHA DE TU MADRE", company)

	return movements, errors.NewBadRequestError("prueba", "22")
}
