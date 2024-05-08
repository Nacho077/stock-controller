package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"net/http"
	"strconv"
)

type DeleteCompanyById struct {
	CompanyRepository repository.CompanyRepositoryInterface
}

func (u DeleteCompanyById) Handle(ctx *gin.Context) {
	companyId, err := strconv.ParseInt(ctx.Param("companyId"), 10, 64)
	if err != nil {
		ctx.JSON(errors.HandleError(errors.NewBadRequestError("Invalid company id", "User Error")))
		return
	}

	err = u.CompanyRepository.DeleteCompanyById(companyId)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
	}

	ctx.JSON(http.StatusOK, "Delete company successfully")
}
