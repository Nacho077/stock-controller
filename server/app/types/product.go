package types

type Product struct {
	Id        *int   `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Brand     string `json:"brand"`
	Detail    string `json:"detail"`
	CompanyId int64  `json:"company_id"`
}
