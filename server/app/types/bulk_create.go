package types

type MovementsDataToSave struct {
	Date         BasicTime `json:"date"`
	ShippingCode string    `json:"shipping_code"`
	Pallets      int       `json:"pallets,omitempty"`
	Units        int       `json:"units,omitempty"`
	Code         string    `json:"code"`
	Name         string    `json:"name"`
	Brand        string    `json:"brand,omitempty"`
	Detail       string    `json:"detail,omitempty"`
	Deposit      string    `json:"deposit,omitempty"`
	Observations string    `json:"observations,omitempty"`
}

type DataToSave struct {
	CompanyName   string                `json:"company_name"`
	MovementsData []MovementsDataToSave `json:"movements_data"`
}
