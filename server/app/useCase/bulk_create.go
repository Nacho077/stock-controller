package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/types"
	"net/http"
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

	return bodyData, nil
}
