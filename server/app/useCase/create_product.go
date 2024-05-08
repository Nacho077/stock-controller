package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/types"
	"net/http"
)

type CreateProduct struct {
	ProductRepository repository.ProductRepositoryInterface
}

func (u CreateProduct) Handle(ctx *gin.Context) {
	var product types.Product

	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	if product.Code == "" {
		ctx.JSON(errors.HandleError(errors.NewBadRequestError("Product Code is required", "User Error")))
		return
	}

	productId, exists, err := u.ProductRepository.CreateProductIfNotExist(product)
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	if exists {
		ctx.JSON(errors.HandleError(errors.NewBadRequestError("Product already exists", "User Error")))
		return
	}

	product.Id = productId
	ctx.JSON(http.StatusOK, product)
}
