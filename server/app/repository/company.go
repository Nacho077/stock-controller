package repository

import (
	"database/sql"
	"fmt"
	"github.com/stock-controller/app/errors"
)

type CompaniesRepositoryI interface {
	GetCompanies() ([]Company, error)
}

type CompaniesRepository struct {
	Db *sql.DB
}

type Company struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (repository *CompaniesRepository) GetCompanies() ([]Company, error) {
	var companies []Company

	rows, err := repository.Db.Query("SELECT * FROM company")
	if err != nil {
		return nil, errors.NewFailedDependency(fmt.Sprintf("Error in database when bringing all companies: %s", err.Error()))
	}

	for rows.Next() {
		var company Company
		if err := rows.Scan(&company.Id, &company.Name); err != nil {
			return nil, errors.NewInternalServerError(fmt.Sprintf("Error in database when converting rows to Companies: %s", err.Error()))
		}
		companies = append(companies, company)
	}

	return companies, nil
}
