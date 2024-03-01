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
	page := ctx.Query("page")
	pageSize := ctx.Query("page_size")
	//codeFilter := ctx.Query("code")
	//nameFilter := ctx.Query("name")
	//brandFilter := ctx.Query("brand")
	orderBy := ctx.Query("order_by")
	whatOrder := ctx.Query("what_order")

	limit, offset, err := repository.getOffset(page, pageSize)

	if err != nil {
		status, errMessage := errors.HandleError(err)
		ctx.JSON(status, errMessage)
		return
	}

	parsedId, err := repository.validateId(id)
	if err != nil {
		status, errMessage := errors.HandleError(err)
		ctx.JSON(status, errMessage)
		return
	}

	filter := ""

	movementsResult, err := repository.MovementRepository.GetMovementsByCompanyId(parsedId, limit, offset, filter, orderBy, whatOrder)
	if err != nil {
		status, errMessage := errors.HandleError(err)
		ctx.JSON(status, errMessage)
		return
	}

	ctx.JSON(http.StatusOK, movementsResult)
}

func (repository GetMovementsByCompany) validateId(id string) (int, error) {
	parsedId, err := strconv.Atoi(id)

	if err != nil {
		return parsedId, errors.NewBadRequestError("Error in Id, id be must a number", err.Error())
	}

	if parsedId < 0 {
		return parsedId, errors.NewBadRequestError("Error in Id, id be must a positive number", err.Error())
	}

	return parsedId, nil
}

func (repository GetMovementsByCompany) getOffset(page string, pageSize string) (int, int, error) {

	if page == "" || page == "0" {
		page = "1"
	}

	if pageSize == "" {
		pageSize = "10"
	}

	parsedPage, err := strconv.Atoi(page)
	if err != nil {
		return parsedPage, 0, errors.NewBadRequestError("Error in Page, page be must a number", err.Error())
	}

	parsedPageSize, err := strconv.Atoi(pageSize)
	if err != nil {
		return parsedPageSize, 0, errors.NewBadRequestError("Error in page size, size be must a number", err.Error())
	}

	offset := (parsedPage - 1) * parsedPageSize

	return parsedPageSize, offset, nil
}

//func (repository GetMovementsByCompany) validateFilters(filter string, order string) () {
//	filter
//}
