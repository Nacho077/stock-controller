package types

type Movement struct {
	Id           *int
	Date         string //time.Time
	ShippingCode string
	Pallets      int
	Units        int
	Deposit      string
	Observations string
}
