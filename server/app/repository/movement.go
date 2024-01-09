package repository

import (
	"fmt"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/types"
)

type MovementRepositoryInterface interface {
	GetMovementsByCompanyId(id int) ([]types.ProductMovement, error)
	CreateMovement(movement types.Movement, productId *int64) error
}

func (repository Repository) GetMovementsByCompanyId(id int) ([]types.ProductMovement, error) {
	var company types.Company

	err := repository.Db.QueryRow("SELECT * FROM company WHERE company.id = ?", id).Scan(&company.Id, &company.Name)
	if err != nil {
		errors.NewFailedDependencyError(fmt.Sprintf("Error in database when bringing company with id %d", id), err.Error())
	}

	if company.Id == 0 {
		return nil, errors.NewBadRequestError(fmt.Sprintf("Company with id %d doesn't exist", id), "User error")
	}

	query := "SELECT product.*, movement.* FROM company"
	query += " INNER JOIN product ON product.company_id = company.id"
	query += " INNER JOIN movements_products ON movements_products.product_id = product.id"
	query += " INNER JOIN movement ON movement.id = movements_products.movement_id"
	query += " WHERE company.id = ?"

	movementsRow, err := repository.Db.Query(query, id)
	if err != nil {
		errors.NewFailedDependencyError("Error in database when bringing movements and related products", err.Error())
	}

	movements := make([]types.ProductMovement, 0)
	var movement types.ProductMovement

	for movementsRow.Next() {
		err = movementsRow.Scan(&movement.ProductId, &movement.Code, &movement.Name, &movement.Brand, &movement.Detail, &movement.CompanyId, &movement.MovementId, &movement.Date, &movement.ShippingCode, &movement.Units, &movement.Deposit, &movement.Observations)
		if err != nil {
			return nil, errors.NewInternalServerError("Error in scan when converting movement", err.Error())
		}
		movements = append(movements, movement)
	}

	return movements, nil
}

func (repository Repository) CreateMovement(movement types.Movement, productId *int64) error {
	// Create new movement

	if productId == nil {
		return errors.NewInternalServerError("Error in Movement when trying to get product id", "Internal Error")
	}

	emptyValues := "?, ?, ?"
	nameValues := "date, shipping_code, units"
	values := []interface{}{movement.Date, movement.ShippingCode, movement.Units}

	if movement.Deposit != nil && *movement.Deposit != "" {
		emptyValues += ", ?"
		nameValues += ", deposit"
		values = append(values, movement.Deposit)
	}

	if movement.Observations != nil && *movement.Observations != "" {
		emptyValues += ", ?"
		nameValues += ", observations"
		values = append(values, movement.Observations)
	}

	query := fmt.Sprintf("INSERT INTO movement(%s) VALUES (%s)", nameValues, emptyValues)

	result, err := repository.Db.Exec(query, values...)

	if err != nil {
		return errors.NewFailedDependencyError("Error when trying to save movements", err.Error())
	}
	movementId, _ := result.LastInsertId()

	// Create relation between products and movements
	if _, err = repository.Db.Exec("INSERT INTO movements_products(movement_id, product_id) VALUES (?, ?)", movementId, productId); err != nil {
		return errors.NewFailedDependencyError("Error when trying to save movements_products", err.Error())
	}

	return nil
}
