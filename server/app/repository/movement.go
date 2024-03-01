package repository

import (
	"fmt"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/types"
	"strings"
)

type MovementRepositoryInterface interface {
	GetMovementsByCompanyId(id int, limit int, offset int, filter string, orderBy string, whatOrder string) (types.MovementsResponse, error)
	CreateMovement(movement types.Movement, productId *int64) error
}

func (repository Repository) GetMovementsByCompanyId(id int, limit int, offset int, filter string, orderBy string, whatOrder string) (types.MovementsResponse, error) {
	var response = types.MovementsResponse{}

	company, err := repository.getCompanyById(id)
	if err != nil {
		return response, err
	}

	response.CompanyName = company.Name

	values := []interface{}{id, limit, offset}
	if filter != "" {
		strings.ToLower(filter)
		values = append(values, filter)
	}

	//if orderBy == "" {
	//	orderBy = "id"
	//}
	//
	//strings.ToLower(orderBy)
	//values = append(values, orderBy)
	//
	//if whatOrder == "" {
	//	whatOrder = "ASC"
	//}
	//
	//strings.ToUpper(whatOrder)
	//values = append(values, whatOrder)

	fmt.Println("TU PUTA MADRE", values)
	query := "SELECT product.*, movement.* FROM company"
	query += " INNER JOIN product ON product.company_id = company.id"
	query += " INNER JOIN movements_products ON movements_products.product_id = product.id"
	query += " INNER JOIN movement ON movement.id = movements_products.movement_id"
	query += " WHERE company.id = ?"
	query += " ORDER BY ? ?"
	query += " LIMIT ? OFFSET ?"

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
