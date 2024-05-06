package types

import (
	"fmt"
	"github.com/stock-controller/app/utils"
	"reflect"
	"strings"
)

type MovementQueries struct {
	CompanyId int64
	MovementFilters
	*Pagination
	Movement
}

func (q MovementQueries) GetQuery() (string, []interface{}) {
	values := []interface{}{q.CompanyId}
	queryFilters := " c.id = ?"
	var orderAndPaginate string

	if q.MovementId != 0 {
		queryFilters += " AND m.Id = ?"
		values = append(values, q.MovementId)
	}

	if q.Code != "" {
		queryFilters += " AND p.code = ?"
		strings.ToLower(q.Code)
		values = append(values, q.Code)
	}

	if q.Brand != "" {
		queryFilters += " AND p.brand = ?"
		strings.ToLower(q.Brand)
		values = append(values, q.Brand)
	}

	if q.Name != "" {
		queryFilters += " AND p.name = ?"
		strings.ToLower(q.Name)
		values = append(values, q.Name)
	}

	if q.Pagination != nil {
		if q.OrderBy == "" {
			q.OrderBy = "id"
		}

		if q.OrderBy != "" {
			movementStruct := reflect.TypeOf(Movement{})
			_, exist := movementStruct.FieldByName(utils.ToCapitalize(q.OrderBy))
			if !exist {
				q.OrderBy = "id"
			}
		}

		if q.OrderDirection == "" || (q.OrderDirection != "ASC" && q.OrderDirection != "DESC") {
			q.OrderDirection = "DESC"
		}

		orderAndPaginate = fmt.Sprintf(" ORDER BY m.%s %s LIMIT ? OFFSET ?", q.OrderBy, q.OrderDirection)
		values = append(values, q.Limit, q.Offset)
	}

	query := "SELECT p.*, m.* FROM company c"
	query += " INNER JOIN product p ON p.company_id = c.id"
	query += " INNER JOIN movements_products mp ON mp.product_id = p.id"
	query += " INNER JOIN movement m ON m.id = mp.movement_id"
	query += " WHERE" + queryFilters
	query += orderAndPaginate

	return query, values
}

func (q MovementQueries) CreateQuery() (string, []interface{}) {
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

func (q MovementQueries) UpdateQuery() (string, []interface{}) {
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
