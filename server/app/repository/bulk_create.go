package repository

import (
	"github.com/stock-controller/app/errors"
	"strings"
	"time"
)

type BulkCreateRepositoryInterface interface {
	BulkCreateData(dataToSave []any) error
}

// VER FECHA

type MovementsDataToSave struct {
	Date         time.Time `json:"date"`
	ShippingCode string    `json:"shipping_code"`
	Pallets      string    `json:"pallets"`
	Units        int       `json:"units"`
	Code         string    `json:"code"`
	Name         string    `json:"name"`
	Brand        string    `json:"brand"`
	Detail       string    `json:"detail"`
	Deposit      string    `json:"deposit"`
	Observations string    `json:"observations"`
}

type DataToSave struct {
	CompanyName   string                `json:"company_name"`
	MovementsData []MovementsDataToSave `json:"movements_data"`
}

func (repository Repository) BulkCreateData(dataToSave DataToSave) error {

	// Create name company
	result, err := repository.Db.Exec("INSERT IGNORE INTO company(name) VALUES ('?')", dataToSave.CompanyName)
	if err != nil {
		errors.NewFailedDependencyError("Error in create company name", err.Error())
	}

	companyId, err := result.LastInsertId()

	// Create all company products
	fieldsProductsEmpty := "('?', '?', '?', '?', '?'), "

	queryProductsBulk := "INSERT IGNORE INTO product(code, name, brand, detail, companyId) VALUES "
	var fieldsProductsValues []interface{}

	for _, data := range dataToSave.MovementsData {
		queryProductsBulk += fieldsProductsEmpty
		fieldsProductsValues = append(fieldsProductsValues, data.Code, data.Name, data.Brand, data.Detail, companyId)
	}

	queryProductsBulk = strings.TrimPrefix(queryProductsBulk, ", ")

	_, err = repository.Db.Query(queryProductsBulk, fieldsProductsValues...)
	if err != nil {
		errors.NewFailedDependencyError("Error in create movements bulk", err.Error())
	}

	// Create all company movements
	queryMovementsBulk := "INSERT IGNORE INTO movement(date, shipping_code, pallets, units,deposit, observations) VALUES "
	fieldsMovementsEmpty := "('?', '?', '?', '?', '?', '?'), "
	var fieldsMovementsValues []interface{}

	for _, data := range dataToSave.MovementsData {
		queryMovementsBulk += fieldsMovementsEmpty
		fieldsMovementsValues = append(fieldsMovementsValues, data.Code, data.ShippingCode, data.Pallets, data.Units, data.Deposit, data.Observations)
	}

	queryMovementsBulk = strings.TrimPrefix(queryMovementsBulk, ", ") + " RETURNING id"

	_, err = repository.Db.Query(queryMovementsBulk, fieldsMovementsValues...)
	if err != nil {
		errors.NewFailedDependencyError("Error in create movements bulk", err.Error())
	}

	return nil
}
