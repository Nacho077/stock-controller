package useCase

import (
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

func (u BulkCreate) Handle(ctx *gin.Context) {
	bodyData, err := u.getBodyData(ctx)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	if err = u.BulkCreateRepository.BulkCreateData(bodyData); err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, "data added successfully")
}

func (u BulkCreate) getBodyData(ctx *gin.Context) (types.DataToSave, error) {
	var bodyData types.DataToSave

	if err := ctx.ShouldBindJSON(&bodyData); err != nil {
		return bodyData, errors.NewInternalServerError("Error in get request body", err.Error())
	}

	bodyData.CompanyName = strings.ToLower(bodyData.CompanyName)

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
