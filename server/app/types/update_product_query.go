package types

import "fmt"

type UpdateProductQuery struct {
	Product
}

func (q UpdateProductQuery) GetQuery() (string, []interface{}) {
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
