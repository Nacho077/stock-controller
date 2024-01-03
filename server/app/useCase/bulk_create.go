package useCase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/types"
	"net/http"
	"strings"
)

type BulkCreate struct {
	BulkCreateRepository repository.BulkCreateRepositoryInterface
}

func (repository BulkCreate) Handle(ctx *gin.Context) {
	bodyData, err := repository.getBodyData(ctx)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	if err = repository.BulkCreateRepository.BulkCreateData(bodyData); err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, "data added successfully")
}

func (repository BulkCreate) getBodyData(ctx *gin.Context) (types.DataToSave, error) {
	var bodyData types.DataToSave

	if err := ctx.ShouldBindJSON(&bodyData); err != nil {
		return bodyData, errors.NewInternalServerError("Error in get request body", err.Error())
	}

	bodyData.CompanyName = strings.ToLower(bodyData.CompanyName)

	fmt.Println("DATE", bodyData.MovementsData[0].Date)

	for i, data := range bodyData.MovementsData {
		bodyData.MovementsData[i].Code = strings.ToLower(data.Code)
		bodyData.MovementsData[i].ShippingCode = strings.ToLower(data.ShippingCode)
		bodyData.MovementsData[i].Name = strings.ToLower(data.Name)
		bodyData.MovementsData[i].Brand = strings.ToLower(data.Brand)
		bodyData.MovementsData[i].Detail = strings.ToLower(data.Detail)
		bodyData.MovementsData[i].Deposit = strings.ToLower(data.Deposit)
		bodyData.MovementsData[i].Observations = strings.ToLower(data.Observations)
	}

	return bodyData, nil
}
