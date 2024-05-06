package repository

import (
	"database/sql"
	goErrors "errors"
	"fmt"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/types"
)

type CompanyRepositoryInterface interface {
	GetCompanies() ([]types.Company, error)
	getCompanyById(id int) (types.Company, error)
	GetCompanyIdByName(name string) (int64, error)
	CreateCompanyIfNotExist(name string) (int64, error)
	DeleteCompanyById(id int64) error
	UpdateCompanyById(company types.Company) error
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

func (repository Repository) getCompanyById(id int) (types.Company, error) {
	var company types.Company
	err := repository.Db.QueryRow("SELECT * FROM company WHERE company.id = ?", id).Scan(&company.Id, &company.Name)
	if err != nil {
		if goErrors.As(err, &sql.ErrNoRows) {
			return company, errors.NewBadRequestError("Company id not exist", err.Error())
		}
		return company, errors.NewFailedDependencyError(fmt.Sprintf("Error in database when bringing company with id %d", id), err.Error())
	}

	if company.Id == nil {
		return company, errors.NewBadRequestError(fmt.Sprintf("Company with id %d doesn't exist", id), "User error")
	}

	return company, nil
}

func (repository Repository) CreateCompanyIfNotExist(name string) (int64, error) {
	var companyId int64

	result, err := repository.Db.Exec("INSERT IGNORE INTO company(name) VALUES (?)", name)
	if err != nil {
		return companyId, errors.NewFailedDependencyError("Error in create company", err.Error())
	}

	companyId, _ = result.LastInsertId()

	return companyId, nil
}

func (repository Repository) GetCompanyIdByName(name string) (int64, error) {
	var companyId int64

	err := repository.Db.QueryRow("SELECT id FROM company WHERE name = ?", name).Scan(&companyId)
	if err != nil {
		return companyId, errors.NewFailedDependencyError("Error in get company id with name", err.Error())
	}

	return companyId, nil
}

func (repository Repository) UpdateCompanyById(company types.Company) error {
	_, err := repository.Db.Exec("UPDATE company SET name = ? WHERE id = ?", company.Name, company.Id)
	if err != nil {
		return errors.NewFailedDependencyError("Error in update company with name", err.Error())
	}

	return nil
}

func (repository Repository) DeleteCompanyById(id int64) error {
	_, err := repository.Db.Exec("DELETE FROM company WHERE id = ?", id)
	if err != nil {
		return errors.NewFailedDependencyError("Error in delete company", err.Error())
	}

	return nil
}
