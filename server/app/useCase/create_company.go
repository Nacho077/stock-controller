package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/types"
	"net/http"
	"regexp"
)

type CreateCompany struct {
	CompanyRepository repository.CompanyRepositoryInterface
}

func (u CreateCompany) Handle(ctx *gin.Context) {
	var company types.Company

	err := ctx.BindJSON(&company)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}
	name := company.Name

	if err = u.isNameValid(name); err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	companyId, err := u.CompanyRepository.CreateCompanyIfNotExist(name)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	if companyId == 0 {
		ctx.JSON(errors.HandleError(errors.NewBadRequestError("Company already exists", "User Error")))
		return
	}
	company.Id = &companyId

	ctx.JSON(http.StatusOK, company)
}

func (u CreateCompany) isNameValid(name string) error {
	matched, _ := regexp.MatchString("[0-9]+", name)
	if matched || name == "" {
		return errors.NewBadRequestError("Company name is not valid", "User Error")
	}

	return nil
}
