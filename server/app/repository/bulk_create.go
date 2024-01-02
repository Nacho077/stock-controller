package repository

import (
	"github.com/stock-controller/app/errors"
	"github.com/stock-controller/app/types"
)

type BulkCreateRepositoryInterface interface {
	BulkCreateData(dataToSave types.DataToSave) error
}

//poner guardar toodo en minúscula

func (repository Repository) BulkCreateData(dataToSave types.DataToSave) error {

	// Create name company
	result, err := repository.Db.Exec("INSERT IGNORE INTO company(name) VALUES (?)", dataToSave.CompanyName)
	if err != nil {
		return errors.NewFailedDependencyError("Error in create company", err.Error())
	}

	companyId, _ := result.LastInsertId()

	if companyId == 0 {
		var id int64

		err = repository.Db.QueryRow("SELECT id FROM company WHERE name = ?", dataToSave.CompanyName).Scan(&id)
		if err != nil {
			return errors.NewFailedDependencyError("Error in get company id with name", err.Error())
		}

		companyId = id
	}

	productsCreated := make(map[string]int64)

	for _, data := range dataToSave.MovementsData {
		// Create all company products
		var productId int64

		if productsCreated[data.Name] != 0 {
			productId = productsCreated[data.Name]

		} else {
			newProduct := types.Product{Name: data.Name, Code: data.Code, Brand: data.Brand, Detail: data.Detail, CompanyId: companyId}

			productId, err = repository.CreateProduct(newProduct)
			if err != nil {
				return err
			}

			productsCreated[data.Name] = productId
		}

		newMovement := types.Movement{Date: data.Date, ShippingCode: data.ShippingCode, Pallets: data.Pallets, Units: data.Units, Deposit: data.Deposit, Observations: data.Observations}
		if err = repository.CreateMovement(newMovement, productId); err != nil {
			return err
		}
	}

	return nil
}
