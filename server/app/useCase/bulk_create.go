package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"net/http"
)

type BulkCreate struct {
	BulkCreateRepository repository.BulkCreateRepositoryInterface
}

func (repository BulkCreate) Handle(ctx *gin.Context, dataToSave any) {

	err := repository.BulkCreateRepository.BulkCreateData(dataToSave)
	if err != nil {
		status, errMessage := errors.HandleError(err)
		ctx.JSON(status, errMessage)
		return
	}

	ctx.JSON(http.StatusOK, "data added successfully")
}
