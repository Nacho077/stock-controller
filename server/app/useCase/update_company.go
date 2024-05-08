package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/types"
	"net/http"
	"strconv"
)

type UpdateCompanyById struct {
	CompanyRepository repository.CompanyRepositoryInterface
}

func (u UpdateCompanyById) Handle(ctx *gin.Context) {
	companyId, err := strconv.ParseInt(ctx.Param("companyId"), 10, 64)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	var company types.Company

	err = ctx.BindJSON(&company)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	company.Id = &companyId

	if company.Name == "" {
		ctx.JSON(errors.HandleError(errors.NewBadRequestError("Field name is required", "User Error")))
		return
	}

	err = u.CompanyRepository.UpdateCompanyById(company)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, company)
}
