package types

type MovementsResponse struct {
	CompanyName string            `json:"company_name"`
	Movements   []ProductMovement `json:"movements"`
}
