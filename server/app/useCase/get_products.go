package useCase

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/types"
	"net/http"
	"strconv"
)

type GetProducts struct {
	ProductRepository repository.ProductRepositoryInterface
}

func (u GetProducts) Handle(ctx *gin.Context) {
	companyId, err := strconv.ParseInt(ctx.Param("companyId"), 10, 64)
	if err != nil {
		ctx.JSON(errors.HandleError(errors.NewBadRequestError("invalid param company id", "User Error")))
		return
	}

	products, err := u.ProductRepository.GetProducts(types.Product{CompanyId: companyId})
	if err != nil {
		ctx.JSON(errors.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, products)
}
