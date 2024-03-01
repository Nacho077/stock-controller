package repository

import (
	"fmt"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/types"
)

type CreateProductRepositoryInterface interface {
	CreateProduct(productToCreate types.Product) (*int64, error)
	CreateProductIfNotExist(productToCreate types.Product) (*int64, error)
	GetProducts(product types.Product) ([]types.Product, error)
}

func (repository Repository) CreateProductIfNotExist(product types.Product) (*int64, error) {

	productFound, err := repository.GetProducts(product)
	if err != nil {
		return nil, err
	}

	if len(productFound) > 1 {
		return nil, errors.NewBadRequestError("There is not enough information to know if the product exists", "User Error")
	}

	if len(productFound) == 1 {
		return productFound[0].Id, nil
	}

	createdProductId, err := repository.CreateProduct(product)
	if err != nil {
		return createdProductId, err
	}

	return createdProductId, nil
}

func (repository Repository) CreateProduct(product types.Product) (*int64, error) {
	values := []interface{}{product.CompanyId}
	emptyValues := "?"
	valueNames := "company_id"

	if product.Name != nil && *product.Name != "" {
		emptyValues += ", ?"
		valueNames += ", name"
		values = append(values, product.Name)
	}

	if product.Code != "" {
		emptyValues += ", ?"
		valueNames += ", code"
		values = append(values, product.Code)
	}

	if product.Brand != nil && *product.Brand != "" {
		emptyValues += ", ?"
		valueNames += ", brand"
		values = append(values, product.Brand)
	}

	if product.Detail != nil && *product.Detail != "" {
		emptyValues += ", ?"
		valueNames += ", detail"
		values = append(values, product.Detail)
	}

	query := fmt.Sprintf("INSERT INTO product(%s) VALUES (%s)", valueNames, emptyValues)

	result, err := repository.Db.Exec(query, values...)
	if err != nil {
		return nil, errors.NewFailedDependencyError("Error when trying to save product", err.Error())
	}

	idCreated, _ := result.LastInsertId()

	return &idCreated, nil
}

func (repository Repository) GetProducts(product types.Product) ([]types.Product, error) {
	conditionValues := []interface{}{product.CompanyId}
	conditions := "company_id = ?"

	if product.Name != nil && *product.Name != "" {
		conditions += " AND name = ?"
		conditionValues = append(conditionValues, product.Name)
	}

	if product.Code != "" {
		conditions += " AND code = ?"
		conditionValues = append(conditionValues, product.Code)
	}

	if product.Brand != nil && *product.Brand != "" {
		conditions += " AND brand = ?"
		conditionValues = append(conditionValues, product.Brand)
	}

	if product.Detail != nil && *product.Detail != "" {
		conditions += " AND detail = ?"
		conditionValues = append(conditionValues, product.Detail)
	}

	query := fmt.Sprintf("SELECT * FROM product WHERE %s", conditions)

	var productsFound []types.Product

	rows, err := repository.Db.Query(query, conditionValues...)
	if err != nil {
		return nil, errors.NewFailedDependencyError("Error in get product", err.Error())
	}

	if rows != nil {
		var product types.Product
		for rows.Next() {
			err = rows.Scan(&product.Id, product.Name, &product.Code, product.Brand, product.Detail, &product.CompanyId)
			if err != nil {
				fmt.Println(product)
				return nil, errors.NewInternalServerError("Error in scan when trying get product", err.Error())
			}

			productsFound = append(productsFound, product)
		}
	}

	return productsFound, nil
}
