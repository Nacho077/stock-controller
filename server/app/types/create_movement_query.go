package types

import "fmt"

type CreateMovementQuery struct {
	Movement
}

func (q CreateMovementQuery) GetQuery() (string, []interface{}) {
	emptyValues := "?, ?, ?"
	nameValues := "date, shipping_code, units"
	values := []interface{}{q.Date, q.ShippingCode, q.Units}

	if q.Deposit != nil && *q.Deposit != "" {
		emptyValues += ", ?"
		nameValues += ", deposit"
		values = append(values, q.Deposit)
	}

	if q.Observations != nil && *q.Observations != "" {
		emptyValues += ", ?"
		nameValues += ", observations"
		values = append(values, q.Observations)
	}

	query := fmt.Sprintf("INSERT INTO movement(%s) VALUES (%s)", nameValues, emptyValues)

	return query, values
}
