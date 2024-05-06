package repository

import (
	"database/sql"
	goErrors "errors"
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
	productQueries := types.ProductQueries{Product: product}
	query, values := productQueries.GetQuery()

	rows, err := repository.Db.Query(query, values...)
	if err != nil {
		return nil, errors.NewFailedDependencyError("Error in get product", err.Error())
	}

	var productsFound []types.Product

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
		if goErrors.As(err, &sql.ErrNoRows) {
			return product, errors.NewBadRequestError("Product id not exist", err.Error())
		}
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
	productQueries := types.ProductQueries{Product: product}
	query, values := productQueries.CreateQuery()

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

	productQueries := types.ProductQueries{Product: product}
	query, values := productQueries.UpdateQuery()

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
