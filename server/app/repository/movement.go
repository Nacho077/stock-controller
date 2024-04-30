package repository

import (
	"fmt"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/types"
)

type MovementRepositoryInterface interface {
	GetMovementsByCompanyId(companyId int, pagination types.Pagination, filters types.MovementFilters) (types.MovementsResponse, error)
	GetMovementById(id int64) (types.Movement, error)
	CreateMovement(movement types.Movement, productId *int64) error
	UpdateMovementById(movement types.Movement) error
	DeleteMovementById(id int64) error
}

func (repository Repository) GetMovementsByCompanyId(companyId int, pagination types.Pagination, filters types.MovementFilters) (types.MovementsResponse, error) {
	var response = types.MovementsResponse{}

	company, err := repository.getCompanyById(companyId)
	if err != nil {
		return response, err
	}

	response.CompanyName = company.Name

	movementQuery := types.MovementQuery{CompanyId: companyId, MovementFilters: filters, Pagination: pagination}
	query, values := movementQuery.GetQuery()

	movementsRow, err := repository.Db.Query(query, values...)
	if err != nil {
		return response, errors.NewFailedDependencyError("Error in database when bringing movements and related products", err.Error())
	}

	movements := make([]types.ProductMovement, 0)
	var movement types.ProductMovement

	if movementsRow == nil {
		return response, nil
	}

	for movementsRow.Next() {
		err = movementsRow.Scan(&movement.ProductId, &movement.Code, &movement.Name, &movement.Brand, &movement.Detail, &movement.CompanyId, &movement.MovementId, &movement.Date, &movement.ShippingCode, &movement.Units, &movement.Deposit, &movement.Observations)
		if err != nil {
			return response, errors.NewInternalServerError("Error in scan when converting movement", err.Error())
		}
		movements = append(movements, movement)
	}

	response.Movements = movements

	return response, nil
}

func (repository Repository) GetMovementById(id int64) (types.Movement, error) {
	var movement types.Movement

	err := repository.Db.QueryRow("SELECT * FROM movement WHERE id = ?", id).Scan(&movement.Id, &movement.Date, &movement.ShippingCode, &movement.Units, &movement.Deposit, &movement.Observations)
	if err != nil {
		return movement, errors.NewFailedDependencyError("Error in get movement by id", err.Error())
	}

	return movement, nil
}

func (repository Repository) CreateMovement(movement types.Movement, productId *int64) error {
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

func (repository Repository) UpdateMovementById(movement types.Movement) error {
	var values []interface{}
	var keys string

	if movement.Date != "" {
		values = append(values, movement.Date)
		keys += "date = ?"
	}

	if movement.ShippingCode != nil {
		values = append(values, movement.ShippingCode)
		if len(keys) > 1 {
			keys += ", "
		}

		keys += "shipping_code = ?"
	}

	if movement.Units != 0 {
		values = append(values, movement.Units)
		if len(keys) > 1 {
			keys += ", "
		}

		keys += "units = ?"
	}

	if movement.Deposit != nil {
		values = append(values, movement.Deposit)
		if len(keys) > 1 {
			keys += ", "
		}

		keys += "deposit = ?"
	}

	if movement.Observations != nil {
		values = append(values, movement.Observations)
		if len(keys) > 1 {
			keys += ", "
		}

		keys += "observations = ?"
	}

	query := fmt.Sprintf("UPDATE movement SET %s WHERE id = ?", keys)
	values = append(values, movement.Id)

	_, err := repository.Db.Exec(query, values...)
	if err != nil {
		return errors.NewFailedDependencyError("Error in update movement by id", err.Error())
	}

	return nil
}

func (repository Repository) DeleteMovementById(id int64) error {
	_, err := repository.Db.Exec("DELETE FROM movements_products WHERE movement_id = ?", id)
	if err != nil {
		return errors.NewFailedDependencyError("Error when trying to delete relation between movement and product", err.Error())
	}

	_, err = repository.Db.Exec("DELETE FROM movement WHERE id = ?", id)
	if err != nil {
		return errors.NewFailedDependencyError("Error when trying to delete movement", err.Error())
	}

	return nil

}
