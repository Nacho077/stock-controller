package types

import "fmt"

type UpdateMovementQuery struct {
	Movement
}

func (q UpdateMovementQuery) GetQuery() (string, []interface{}) {
	var keys string
	var values []interface{}

	if q.Date != "" {
		values = append(values, q.Date)
		keys += "date = ?"
	}

	if q.ShippingCode != nil {
		values = append(values, q.ShippingCode)
		if len(keys) > 1 {
			keys += ", "
		}

		keys += "shipping_code = ?"
	}

	if q.Units != 0 {
		values = append(values, q.Units)
		if len(keys) > 1 {
			keys += ", "
		}

		keys += "units = ?"
	}

	if q.Deposit != nil {
		values = append(values, q.Deposit)
		if len(keys) > 1 {
			keys += ", "
		}

		keys += "deposit = ?"
	}

	if q.Observations != nil {
		values = append(values, q.Observations)
		if len(keys) > 1 {
			keys += ", "
		}

		keys += "observations = ?"
	}

	query := fmt.Sprintf("UPDATE movement SET %s WHERE id = ?", keys)
	values = append(values, q.Id)

	return query, values
}
