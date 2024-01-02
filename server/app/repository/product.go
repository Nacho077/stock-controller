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

	repository.Db.QueryRow("SELECT id FROM product WHERE name = ? AND code = ? AND  company_id = ?", productToCreate.Name, productToCreate.Code, productToCreate.CompanyId).Scan(&productId)

	if productId != 0 {
		return productId, nil
	}

	var values []interface{}

	emptyValues := "(?, ?, ?)"
	valueNames := "code, name, company_id"
	values = append(values, productToCreate.Code, productToCreate.Name, productToCreate.CompanyId)

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

	query := fmt.Sprintf("INSERT INTO product(%s) VALUES %s", valueNames, emptyValues)

	result, err := repository.Db.Exec(query, values...)
	if err != nil {
		return productId, errors.NewFailedDependencyError("Error when trying to save product", err.Error())
	}

	return result.LastInsertId()
}
