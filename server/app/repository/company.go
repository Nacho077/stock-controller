package repository

import (
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/types"
)

type CompaniesRepositoryInterface interface {
	GetCompanies() ([]types.Company, error)
}

func (repository Repository) GetCompanies() ([]types.Company, error) {
	companies := make([]types.Company, 0)

	rows, err := repository.Db.Query("SELECT * FROM company")
	if err != nil {
		return nil, errors.NewFailedDependencyError("Error in database when bringing all companies", err.Error())
	}

	for rows.Next() {
		var company types.Company
		if err := rows.Scan(&company.Id, &company.Name); err != nil {
			return nil, errors.NewInternalServerError("Error in database when converting rows to Companies", err.Error())
		}
		companies = append(companies, company)
	}

	return companies, nil
}
