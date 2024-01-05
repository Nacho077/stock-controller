package repository

import (
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/types"
)

type CompaniesRepositoryInterface interface {
	GetCompanies() ([]types.Company, error)
	CreateCompany(name string) error
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

func (repository Repository) CreateCompany(name string) (int64, error) {
	var companyId int64

	result, err := repository.Db.Exec("INSERT IGNORE INTO company(name) VALUES (?)", name)
	if err != nil {
		return companyId, errors.NewFailedDependencyError("Error in create company", err.Error())
	}

	companyId, _ = result.LastInsertId()

	if companyId == 0 {
		var id int64

		err = repository.Db.QueryRow("SELECT id FROM company WHERE name = ?", name).Scan(&id)
		if err != nil {
			return companyId, errors.NewFailedDependencyError("Error in get company id with name", err.Error())
		}

		companyId = id
	}

	return companyId, nil
}
