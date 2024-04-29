package types

import "fmt"

type CreateProductQuery struct {
	Product
}

func (q CreateProductQuery) GetQuery() (string, []interface{}) {
	values := []interface{}{q.CompanyId}
	emptyValues := "?"
	valueNames := "company_id"

	if q.Name != nil && *q.Name != "" {
		emptyValues += ", ?"
		valueNames += ", name"
		values = append(values, q.Name)
	}

	if q.Code != "" {
		emptyValues += ", ?"
		valueNames += ", code"
		values = append(values, q.Code)
	}

	if q.Brand != nil && *q.Brand != "" {
		emptyValues += ", ?"
		valueNames += ", brand"
		values = append(values, q.Brand)
	}

	if q.Detail != nil && *q.Detail != "" {
		emptyValues += ", ?"
		valueNames += ", detail"
		values = append(values, q.Detail)
	}

	query := fmt.Sprintf("INSERT INTO product(%s) VALUES (%s)", valueNames, emptyValues)

	return query, values
}
