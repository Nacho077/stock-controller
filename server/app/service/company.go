package service

import (
	"github.com/stock-controller/app/repository"
)

type CompanyService struct {
	CompanyRepository repository.CompaniesRepositoryI
}

func (service *CompanyService) GetCompanies() ([]repository.Company, error) {
	return service.CompanyRepository.GetCompanies()
}
