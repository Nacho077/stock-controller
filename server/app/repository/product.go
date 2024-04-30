package repository

import (
	"fmt"
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/types"
)

type ProductRepositoryInterface interface {
	GetProducts(product types.Product) ([]types.Product, error)
	GetProductById(productId int64) (types.Product, error)
	CreateProduct(product types.Product) (*int64, error)
	CreateProductIfNotExist(product types.Product) (*int64, bool, error)
	UpdateProductById(product types.Product) error
	DeleteProductById(id int64) error
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
			err = rows.Scan(&product.Id, &product.Name, &product.Code, &product.Brand, &product.Detail, &product.CompanyId)
			if err != nil {
				return nil, errors.NewInternalServerError("Error in scan when trying get product", err.Error())
			}

			productsFound = append(productsFound, product)
		}
	}

	return productsFound, nil
}

func (repository Repository) GetProductById(id int64) (types.Product, error) {
	var product types.Product

	err := repository.Db.QueryRow("SELECT * FROM product WHERE id = ?", id).Scan(&product.Id, &product.Code, &product.Name, &product.Brand, &product.Detail, &product.CompanyId)
	if err != nil {
		return product, errors.NewFailedDependencyError("Error in get product by id", err.Error())
	}

	return product, nil
}

func (repository Repository) CreateProductIfNotExist(product types.Product) (*int64, bool, error) {

	productFound, err := repository.GetProducts(product)
	if err != nil {
		return nil, false, err
	}

	if len(productFound) > 1 {
		return nil, false, errors.NewBadRequestError("There is not enough information to know if the product exists", "User Error")
	}

	if len(productFound) == 1 {
		return productFound[0].Id, true, nil
	}

	createdProductId, err := repository.CreateProduct(product)
	if err != nil {
		return createdProductId, false, err
	}

	return createdProductId, false, nil
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

func (repository Repository) UpdateProductById(product types.Product) error {
	if product.Code == "" {
		return errors.NewBadRequestError("Field code is required", "User Error")
	}

	values := []interface{}{product.Code}
	keys := "code = ?"

	if product.Name != nil {
		values = append(values, product.Name)
		keys += ", name = ?"
	}

	if product.Brand != nil {
		values = append(values, product.Brand)
		keys += ", brand = ?"
	}

	if product.Detail != nil {
		values = append(values, product.Detail)
		keys += ", detail = ?"
	}

	query := fmt.Sprintf("UPDATE product SET %s WHERE id = ?", keys)
	values = append(values, product.Id)
	fmt.Println("QUERY: ", query)
	_, err := repository.Db.Exec(query, values...)
	if err != nil {
		return errors.NewFailedDependencyError("Error in update product by id", err.Error())
	}

	return nil
}

func (repository Repository) DeleteProductById(id int64) error {
	_, err := repository.Db.Exec("DELETE FROM product WHERE id = ?", id)
	if err != nil {
		return errors.NewFailedDependencyError("Error in deleting product", err.Error())
	}

	return nil
}
