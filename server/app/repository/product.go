package repository

import (
	"fmt"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/types"
	"github.com/stock-controller/app/utils"
)

type CreateProductRepositoryInterface interface {
	CreateProduct(productToCreate types.Product) (int64, error)
}

func (repository Repository) CreateProduct(productToCreate types.Product) (int64, error) {
	var productId int64
	var conditionValues []interface{}

	conditions := "company_id = ?"
	conditionValues = append(conditionValues, productToCreate.CompanyId)

	if productToCreate.Name != "" {
		conditions += "name = ?"
		conditionValues = append(conditionValues, productToCreate.Name)
	}

	if productToCreate.Code != "" {
		conditions += "code = ?"
		conditionValues = append(conditionValues, productToCreate.Code)
	}

	if productToCreate.Brand != "" {
		conditions += "brand = ?"
		conditionValues = append(conditionValues, productToCreate.Brand)
	}

	if productToCreate.Detail != "" {
		conditions += "detail = ?"
		conditionValues = append(conditionValues, productToCreate.Detail)
	}

	query := fmt.Sprintf("SELECT id FROM product WHERE %s", conditions)

	repository.Db.QueryRow(query, conditionValues...).Scan(&productId)

	if productId != 0 {
		return productId, nil
	}

	var values []interface{}

	emptyValues := "(?)"
	valueNames := "company_id"
	values = append(values, productToCreate.CompanyId)

	if productToCreate.Name != "" {
		emptyValues = utils.TrimSuffixAndAddText(emptyValues, ")", ", ?)")
		valueNames = utils.TrimSuffixAndAddText(valueNames, ")", ", name")
		values = append(values, productToCreate.Name)
	}

	if productToCreate.Code != "" {
		emptyValues = utils.TrimSuffixAndAddText(emptyValues, ")", ", ?)")
		valueNames = utils.TrimSuffixAndAddText(valueNames, ")", ", code")
		values = append(values, productToCreate.Code)
	}

	if productToCreate.Brand != "" {
		emptyValues = utils.TrimSuffixAndAddText(emptyValues, ")", ", ?)")
		valueNames = utils.TrimSuffixAndAddText(valueNames, ")", ", brand")
		values = append(values, productToCreate.Brand)
	}

	if productToCreate.Detail != "" {
		emptyValues = utils.TrimSuffixAndAddText(emptyValues, ")", ", ?)")
		valueNames = utils.TrimSuffixAndAddText(valueNames, ")", ", detail")
		values = append(values, productToCreate.Detail)
	}

	query = fmt.Sprintf("INSERT INTO product(%s) VALUES %s", valueNames, emptyValues)

	result, err := repository.Db.Exec(query, values...)
	if err != nil {
		return productId, errors.NewFailedDependencyError("Error when trying to save product", err.Error())
	}

	return result.LastInsertId()
}
