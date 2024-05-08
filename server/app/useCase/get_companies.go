package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"net/http"
)

type GetCompanies struct {
	CompanyRepository repository.CompanyRepositoryInterface
}

func (u GetCompanies) Handle(ctx *gin.Context) {
	companies, err := u.CompanyRepository.GetCompanies()
	if err != nil {
		status, errMessage := errors.HandleError(err)
		ctx.JSON(status, errMessage)
		return
	}

	ctx.JSON(http.StatusOK, companies)
}
