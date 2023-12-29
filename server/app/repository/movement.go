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

	row, err := repository.Db.Query("SELECT * FROM company WHERE company.id = ?", id)
	if err != nil {
		errors.NewFailedDependencyError(fmt.Sprintf("Error in database when bringing company with id %d", id), err.Error())
	}

	var company types.Company

	if err := row.Scan(&company.Id, &company.Name); err != nil {
		return nil, errors.NewInternalServerError("Error in scan when converting company", err.Error())
	}

	if company.Id == 0 {
		return nil, errors.NewBadRequestError(fmt.Sprintf("Company with id %d doesn't exist", id), "User error")
	}

	getCompany := "SELECT * FROM company WHERE company.id = ?"
	asociatedProducts := getCompany + "INNER JOIN product ON product.company_id = company.id"
	movementsAsociated := asociatedProducts + "INNER JOIN Movements_products ON Movements_products.product_id = product.id"
	movementsDetail := movementsAsociated + "INNER JOIN Movements ON Movements.id = Movements_products.movement_id"
	selectCompanyById := movementsDetail + "WHERE company.id = 1"

	repository.Db.Query(getCompany + asociatedProducts + movementsAsociated + movementsDetail + selectCompanyById)

	//if company == nil {
	//	errors.NewBadRequestError(fmt.Sprintf("Error when searching for a company, the company with id %d doesn't exist", id), "User error")
	//}

	return movements, nil
}
