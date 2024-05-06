package types

type MovementsResponse struct {
	CompanyName string            `json:"company_name"`
	TotalUnits  int               `json:"total_units"`
	Movements   []ProductMovement `json:"movements"`
}
