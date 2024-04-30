package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/types"
	"net/http"
	"strconv"
)

type UpdateProductById struct {
	ProductRepository repository.ProductRepositoryInterface
}

func (this UpdateProductById) Handle(ctx *gin.Context) {
	productId, err := strconv.ParseInt(ctx.Param("productId"), 10, 64)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	var product types.Product

	err = ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	product.Id = &productId

	err = this.ProductRepository.UpdateProductById(product)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	productUpdated, err := this.ProductRepository.GetProductById(productId)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, productUpdated)
}
