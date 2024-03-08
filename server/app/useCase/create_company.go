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
	CompanyRepository repository.CompaniesRepositoryInterface
}

func (repository CreateCompany) Handle(ctx *gin.Context) {
	var requestBody types.Company
	var companyIdCreated types.CompanyIdCreated

	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}
	name := requestBody.Name

	if err = isNameValid(name); err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	companyId, err := repository.CompanyRepository.CreateCompanyIfNotExist(name)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	if companyId == 0 {
		ctx.JSON(errors.HandleError(errors.NewBadRequestError("Company already exists", "User Error")))
		return
	}
	companyIdCreated.Id = companyId

	ctx.JSON(http.StatusOK, companyIdCreated)
}

func isNameValid(name string) error {
	matched, _ := regexp.MatchString("^[0-9]+$", name)
	if matched || name == "" {
		return errors.NewBadRequestError("Company name is not valid", "User Error")
	}

	return nil
}
