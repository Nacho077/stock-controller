package types

type Movement struct {
	Id           *int64    `json:"id"`
	Date         BasicTime `json:"date"`
	ShippingCode *string   `json:"shipping_code"`
	Units        int       `json:"units"`
	Deposit      *string   `json:"deposit"`
	Observations *string   `json:"observations"`
}
