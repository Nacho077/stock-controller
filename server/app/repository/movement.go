package repository

import (
	"fmt"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/types"
)

type MovementRepositoryInterface interface {
	GetMovementsByCompany(id int) ([]types.Movement, error)
	CreateMovement(movement types.Movement, productId int64) error
}

func (repository Repository) GetMovementsByCompany(id int) ([]types.Movement, error) {

	companyRow, err := repository.Db.Query("SELECT * FROM company WHERE company.id = ?", id)
	if err != nil {
		errors.NewFailedDependencyError(fmt.Sprintf("Error in database when bringing company with id %d", id), err.Error())
	}

	var company types.Company
	if err = companyRow.Scan(&company.Id, &company.Name); err != nil {
		return nil, errors.NewInternalServerError("Error in scan when converting company", err.Error())
	}

	if company.Id == 0 {
		return nil, errors.NewBadRequestError(fmt.Sprintf("Company with id %d doesn't exist", id), "User error")
	}

	query := "SELECT * FROM company"
	query += " INNER JOIN product ON product.company_id = company.id"
	query += " INNER JOIN movements_products ON movements_products.product_id = product.id"
	query += " INNER JOIN movements ON movements.id = movements_products.movement_id"
	query += " WHERE company.id = ?"

	movementsRow, err := repository.Db.Query(query)

	movements := make([]types.Movement, 0)
	var movement types.Movement

	for movementsRow.Next() {
		err = movementsRow.Scan(&movement.Id, &movement.Date, &movement.ShippingCode, &movement.Pallets, &movement.Units, &movement.Deposit, &movement.Observations)
		if err != nil {
			return nil, errors.NewInternalServerError("Error in scan when converting movement", err.Error())
		}
		movements = append(movements, movement)
	}

	return movements, nil
}

func (repository Repository) CreateMovement(movement types.Movement, productId int64) error {
	// Create new movement
	result, err := repository.Db.Exec("INSERT INTO movement(date, shipping_code, pallets, units, deposit, observations) VALUES (?, ?, ?, ?, ?, ?)", movement.Date, movement.ShippingCode, movement.Pallets, movement.Units, movement.Deposit, movement.Observations)
	if err != nil {
		return errors.NewFailedDependencyError("Error when trying to save movements", err.Error())
	}
	movementId, _ := result.LastInsertId()

	// Create relation between products and movements
	if _, err = repository.Db.Exec("INSERT INTO movements_products(movement_id, product_id) VALUES (?, ?)", movementId, productId); err != nil {
		return errors.NewFailedDependencyError("Error when trying to save movements", err.Error())
	}

	return nil
}
