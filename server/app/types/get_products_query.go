package types

import "fmt"

type GetProductsQuery struct {
	Product
}

func (q GetProductsQuery) GetQuery() (string, []interface{}) {
	values := []interface{}{q.CompanyId}
	conditions := "company_id = ?"

	if q.Name != nil && *q.Name != "" {
		conditions += " AND name = ?"
		values = append(values, q.Name)
	}

	if q.Code != "" {
		conditions += " AND code = ?"
		values = append(values, q.Code)
	}

	if q.Brand != nil && *q.Brand != "" {
		conditions += " AND brand = ?"
		values = append(values, q.Brand)
	}

	if q.Detail != nil && *q.Detail != "" {
		conditions += " AND detail = ?"
		values = append(values, q.Detail)
	}

	query := fmt.Sprintf("SELECT * FROM product WHERE %s", conditions)

	return query, values
}
