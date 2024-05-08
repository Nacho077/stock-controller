package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"net/http"
	"strconv"
)

type DeleteMovementById struct {
	MovementRepository repository.MovementRepositoryInterface
}

func (u DeleteMovementById) Handle(ctx *gin.Context) {
	movementId, err := strconv.ParseInt(ctx.Param("movementId"), 10, 64)
	if err != nil {
		ctx.JSON(errors.HandleError(errors.NewBadRequestError("Invalid movement id", "User Error")))
		return
	}

	err = u.MovementRepository.DeleteMovementById(movementId)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, "Deleted movement successfully")
}
