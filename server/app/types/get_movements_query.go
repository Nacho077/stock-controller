package types

import (
	"fmt"
	"strings"
)

type MovementQuery struct {
	CompanyId int
	MovementFilters
	Pagination
}

func (q *MovementQuery) GetQuery() (string, []interface{}) {
	values := []interface{}{q.CompanyId}

	if q.OrderBy == "" {
		q.OrderBy = "id"
	}
	strings.ToLower(q.OrderBy)

	if q.OrderDirection == "" {
		q.OrderDirection = "ASC"
	}
	strings.ToUpper(q.OrderDirection)

	var order = fmt.Sprintf("%s %s", q.OrderBy, q.OrderDirection)
	queryFilters := " c.id = ?"

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

	values = append(values, q.Limit, q.Offset)

	query := "SELECT p.*, m.* FROM company c"
	query += " INNER JOIN product p ON p.company_id = c.id"
	query += " INNER JOIN movements_products mp ON mp.product_id = p.id"
	query += " INNER JOIN movement m ON m.id = mp.movement_id"
	query += " WHERE" + queryFilters
	query += " ORDER BY m." + order
	query += " LIMIT ? OFFSET ?"

	return query, values
}
