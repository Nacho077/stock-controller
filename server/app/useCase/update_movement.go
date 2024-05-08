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

func (u UpdateMovementById) Handle(ctx *gin.Context) {
	movementId, err := strconv.ParseInt(ctx.Param("movementId"), 10, 64)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	var movement types.Movement

	err = ctx.BindJSON(&movement)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	movement.Id = &movementId

	err = u.MovementRepository.UpdateMovementById(movement)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	movementUpdated, err := u.MovementRepository.GetMovementById(movementId)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, movementUpdated)
}
