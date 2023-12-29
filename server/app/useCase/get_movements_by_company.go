package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"net/http"
	"strconv"
)

type GetMovementsByCompany struct {
	MovementRepository repository.MovementRepositoryInterface
}

func (repository GetMovementsByCompany) Handle(ctx *gin.Context) {

	id := ctx.Param("id")

	parsedId, err := repository.validateId(id)
	if err != nil {
		status, errMessage := errors.HandleError(err)
		ctx.JSON(status, errMessage)
		return
	}

	movementsResult, err := repository.MovementRepository.GetMovementsByCompany(parsedId)
	if err != nil {
		status, errMessage := errors.HandleError(err)
		ctx.JSON(status, errMessage)
		return
	}

	ctx.JSON(http.StatusOK, movementsResult)
}

func (controller GetMovementsByCompany) validateId(id string) (int, error) {
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return parsedId, errors.NewBadRequestError("Error in Id, id be must a number", err.Error())
	}

	if parsedId < 0 {
		return parsedId, errors.NewBadRequestError("Error in Id, id be must a positive number", err.Error())
	}

	return parsedId, nil
}
