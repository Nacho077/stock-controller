package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/service"
	"net/http"
)

type CompanyController struct {
	CompanyService service.CompanyService
}

func (controller *CompanyController) GetAllCompanies(ctx *gin.Context) {

	companies, err := controller.CompanyService.GetCompanies()
	if err != nil {
		status, errMessage := errors.HandleError(err)
		ctx.JSON(status, errMessage)
		return
	}

	ctx.JSON(http.StatusOK, companies)
}
