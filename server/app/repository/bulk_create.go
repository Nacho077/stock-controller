package repository

import (
	"github.com/stock-controller/app/types"
)

type BulkCreateRepositoryInterface interface {
	BulkCreateData(dataToSave types.DataToSave) error
}

func (repository Repository) BulkCreateData(dataToSave types.DataToSave) error {

	// Create company
	companyId, err := repository.CreateCompanyIfNotExist(dataToSave.CompanyName)
	if err != nil {
		return err
	}

	if companyId == 0 {
		if companyId, err = repository.GetCompanyId(dataToSave.CompanyName); err != nil {
			return err
		}
	}

	productsCreated := make(map[string]int64)

	for _, data := range dataToSave.MovementsData {
		// Create all company products
		var productId int64

		if productsCreated[data.Name] != 0 {
			productId = productsCreated[data.Name]

		} else {
			newProduct := types.Product{Name: data.Name, Code: data.Code, Brand: data.Brand, Detail: data.Detail, CompanyId: companyId}

			productId, err := repository.CreateProduct(newProduct)
			if err != nil {
				return err
			}

			productsCreated[data.Name] = productId
		}

		newMovement := types.Movement{Date: data.Date, ShippingCode: data.ShippingCode, Units: data.Units, Deposit: data.Deposit, Observations: data.Observations}
		if err := repository.CreateMovement(newMovement, productId); err != nil {
			return err
		}
	}

	return nil
}
