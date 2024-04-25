package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/types"
	"net/http"
	"strconv"
)

type CreateMovement struct {
	MovementRepository repository.MovementRepositoryInterface
}

func (this CreateMovement) Handle(ctx *gin.Context) {
	var movementWithProductId types.MovementWithProductId

	_, err := strconv.ParseInt(ctx.Param("companyId"), 10, 64)
	if err != nil {
		ctx.JSON(errors.HandleError(errors.NewBadRequestError("invalid param company id", "User Error")))
		return
	}

	// Coming soon validations with companyId

	err = ctx.BindJSON(&movementWithProductId)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	err = this.MovementRepository.CreateMovement(movementWithProductId.Movement, &movementWithProductId.ProductId)

	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, "Created Successfully")
}
