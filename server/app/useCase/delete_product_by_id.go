package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"net/http"
	"strconv"
)

type DeleteProductById struct {
	ProductRepository repository.ProductRepositoryInterface
}

func (this DeleteProductById) Handle(ctx *gin.Context) {
	productId, err := strconv.ParseInt(ctx.Param("productId"), 10, 64)
	if err != nil {
		ctx.JSON(errors.HandleError(errors.NewBadRequestError("Invalid product id", "User Error")))
		return
	}

	err = this.ProductRepository.DeleteProductById(productId)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
	}

	ctx.JSON(http.StatusOK, "Delete Product Success")
}
