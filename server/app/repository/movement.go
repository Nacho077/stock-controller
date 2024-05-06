package repository

import (
	"database/sql"
	goErrors "errors"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/types"
)

type MovementRepositoryInterface interface {
	GetMovementsByCompanyId(companyId int64, pagination *types.Pagination, filters types.MovementFilters) (types.MovementsResponse, error)
	GetMovementById(id int64) (types.Movement, error)
	CreateMovement(movement types.Movement, productId *int64) (int64, error)
	UpdateMovementById(movement types.Movement) error
	DeleteMovementById(id int64) error
}

func (repository Repository) GetMovementsByCompanyId(companyId int64, pagination *types.Pagination, filters types.MovementFilters) (types.MovementsResponse, error) {
	var response = types.MovementsResponse{}

	var (
		companyNameChannel = make(chan types.AsyncResponse[string])
		movementsChannel   = make(chan types.AsyncResponse[[]types.ProductMovement])
		totalUnitsChannel  = make(chan types.AsyncResponse[int])
	)

	movementQuery := types.MovementQueries{CompanyId: companyId, MovementFilters: filters, Pagination: pagination}

	go repository.getCompanyName(companyId, companyNameChannel)
	go repository.getMovements(movementQuery, movementsChannel)
	go repository.getTotalUnits(movementQuery, totalUnitsChannel)

	companyNameResponse := <-companyNameChannel
	if companyNameResponse.Error != nil {
		return response, companyNameResponse.Error
	}
	response.CompanyName = companyNameResponse.Data

	movementsResponse := <-movementsChannel
	if movementsResponse.Error != nil {
		return response, movementsResponse.Error
	}
	response.Movements = movementsResponse.Data

	totalUnitsResponse := <-totalUnitsChannel
	if totalUnitsResponse.Error != nil {
		return response, totalUnitsResponse.Error
	}
	response.TotalUnits = totalUnitsResponse.Data

	return response, nil
}

func (repository Repository) getCompanyName(companyId int64, companyNameChannel chan types.AsyncResponse[string]) {
	company, err := repository.getCompanyById(companyId)
	if err != nil {
		companyNameChannel <- types.AsyncResponse[string]{Error: err}
	}

	companyNameChannel <- types.AsyncResponse[string]{company.Name, nil}
}

func (repository Repository) getMovements(movementQuery types.MovementQueries, movementsChannel chan types.AsyncResponse[[]types.ProductMovement]) {
	query, values := movementQuery.GetQuery()

	movementsRow, err := repository.Db.Query(query, values...)
	if err != nil {
		movementsChannel <- types.AsyncResponse[[]types.ProductMovement]{Error: errors.NewFailedDependencyError("Error in database when bringing movements and related products", err.Error())}
	}

	movements := make([]types.ProductMovement, 0)
	var movement types.ProductMovement

	for movementsRow != nil && movementsRow.Next() {
		err = movementsRow.Scan(&movement.ProductId, &movement.Code, &movement.Name, &movement.Brand, &movement.Detail, &movement.CompanyId, &movement.MovementId, &movement.Date, &movement.ShippingCode, &movement.Units, &movement.Deposit, &movement.Observations)
		if err != nil {
			movementsChannel <- types.AsyncResponse[[]types.ProductMovement]{Error: errors.NewInternalServerError("Error in scan when converting movement", err.Error())}
		}
		movements = append(movements, movement)
	}
	movementsChannel <- types.AsyncResponse[[]types.ProductMovement]{Data: movements, Error: nil}
}

func (repository Repository) getTotalUnits(movementQuery types.MovementQueries, totalUnitsChannel chan types.AsyncResponse[int]) {
	unitsQuery, unitsQueryValues := movementQuery.GetTotalUnitsQuery()

	var totalUnits int
	err := repository.Db.QueryRow(unitsQuery, unitsQueryValues...).Scan(&totalUnits)
	if err != nil {
		totalUnitsChannel <- types.AsyncResponse[int]{Error: errors.NewFailedDependencyError("Error in database when bringing total movements units", err.Error())}
	}

	totalUnitsChannel <- types.AsyncResponse[int]{Data: totalUnits, Error: nil}
}

func (repository Repository) GetMovementById(id int64) (types.Movement, error) {
	var movement types.Movement

	err := repository.Db.QueryRow("SELECT * FROM movement WHERE id = ?", id).Scan(&movement.Id, &movement.Date, &movement.ShippingCode, &movement.Units, &movement.Deposit, &movement.Observations)
	if err != nil {
		if goErrors.As(err, &sql.ErrNoRows) {
			return movement, errors.NewBadRequestError("Movement id not exist", err.Error())
		}
		return movement, errors.NewFailedDependencyError("Error in get movement by id", err.Error())
	}

	return movement, nil
}

func (repository Repository) CreateMovement(movement types.Movement, productId *int64) (int64, error) {
	if productId == nil {
		return 0, errors.NewInternalServerError("Error in Movement when trying to get product id", "Internal Error")
	}

	movementQueries := types.MovementQueries{Movement: movement}
	query, values := movementQueries.CreateQuery()

	result, err := repository.Db.Exec(query, values...)

	if err != nil {
		return 0, errors.NewFailedDependencyError("Error when trying to save movements", err.Error())
	}
	movementId, _ := result.LastInsertId()

	// Create relation between products and movements
	if _, err = repository.Db.Exec("INSERT INTO movements_products(movement_id, product_id) VALUES (?, ?)", movementId, productId); err != nil {
		return 0, errors.NewFailedDependencyError("Error when trying to save movements_products", err.Error())
	}

	return movementId, nil
}

func (repository Repository) UpdateMovementById(movement types.Movement) error {
	movementQueries := types.MovementQueries{Movement: movement}
	query, values := movementQueries.UpdateQuery()

	if len(values) == 1 {
		return errors.NewBadRequestError("You must send at least one field to modify", "User Error")
	}

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
