package types

type ProductMovement struct {
	MovementId   int       `json:"movement_id"`
	Date         BasicTime `json:"date"`
	ShippingCode string    `json:"shipping_code"`
	Units        int       `json:"units"`
	Deposit      *string   `json:"deposit"`
	Observations *string   `json:"observations"`
	ProductId    int       `json:"product_id"`
	Name         *string   `json:"name"`
	Code         *string   `json:"code"`
	Brand        *string   `json:"brand"`
	Detail       *string   `json:"detail"`
	CompanyId    int64     `json:"company_id"`
}
