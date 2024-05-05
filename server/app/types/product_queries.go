package types

import "fmt"

type ProductQueries struct {
	Product
}

func (q ProductQueries) GetQuery() (string, []interface{}) {
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

func (q ProductQueries) CreateQuery() (string, []interface{}) {
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

func (q ProductQueries) UpdateQuery() (string, []interface{}) {
	values := []interface{}{q.Code}
	keys := "code = ?"

	if q.Name != nil {
		values = append(values, q.Name)
		keys += ", name = ?"
	}

	if q.Brand != nil {
		values = append(values, q.Brand)
		keys += ", brand = ?"
	}

	if q.Detail != nil {
		values = append(values, q.Detail)
		keys += ", detail = ?"
	}

	query := fmt.Sprintf("UPDATE product SET %s WHERE id = ?", keys)
	values = append(values, q.Id)

	return query, values
}
