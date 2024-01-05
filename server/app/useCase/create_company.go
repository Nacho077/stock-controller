package useCase

import (
	"github.com/stock-controller/app/repository"
)

type CreateCompany struct {
	CompanyRepository repository.CompaniesRepositoryInterface
}

func (repository CreateCompany) Handle(name string) {
	//_, err := repository.CompanyRepository.CreateCompany(name)
}
