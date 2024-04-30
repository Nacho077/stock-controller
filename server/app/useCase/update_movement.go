package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/types"
	"net/http"
	"strconv"
)

type UpdateMovementById struct {
	MovementRepository repository.MovementRepositoryInterface
}

func (this UpdateMovementById) Handle(ctx *gin.Context) {
	movementId, err := strconv.ParseInt(ctx.Param("movementId"), 10, 64)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	movement := types.Movement{Id: &movementId}

	err = ctx.BindJSON(&movement)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	err = this.MovementRepository.UpdateMovementById(movement)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	movementUpdated, err := this.MovementRepository.GetMovementById(movementId)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, movementUpdated)
}
