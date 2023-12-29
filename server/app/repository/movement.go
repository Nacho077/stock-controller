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
	movements := make([]types.Movement, 0)

	companyRow, err := repository.Db.Query("SELECT * FROM company WHERE company.id = ?", id)
	if err != nil {
		errors.NewFailedDependencyError(fmt.Sprintf("Error in database when bringing company with id %d", id), err.Error())
	}

	var company types.Company
	if err := companyRow.Scan(&company.Id, &company.Name); err != nil {
		return nil, errors.NewInternalServerError("Error in scan when converting company", err.Error())
	}

	if company.Id == 0 {
		return nil, errors.NewBadRequestError(fmt.Sprintf("Company with id %d doesn't exist", id), "User error")
	}

	getCompany := "SELECT * FROM company WHERE company.id = ?"
	asociatedProducts := getCompany + " INNER JOIN product ON product.company_id = company.id"
	movementsAsociated := asociatedProducts + " INNER JOIN movements_products ON movements_products.product_id = product.id"
	movementsDetail := movementsAsociated + " INNER JOIN movements ON movements.id = movements_products.movement_id"
	selectCompanyById := movementsDetail + " WHERE company.id = 1"

	movementsRow, err := repository.Db.Query(getCompany + asociatedProducts + movementsAsociated + movementsDetail + selectCompanyById)

	var movement types.Movement
	for movementsRow.Next() {
		movementErr := movementsRow.Scan(&movement.Id, &movement.Date, &movement.ShippingCode, &movement.Pallets, &movement.Units, &movement.Deposit, &movement.Observations)
		if movementErr != nil {
			return nil, errors.NewInternalServerError("Error in scan when converting movement", err.Error())
		}
		movements = append(movements, movement)
	}

	return movements, nil
}
