package types

type Movement struct {
	Id           *int
	Date         BasicTime
	ShippingCode string
	Pallets      int
	Units        int
	Deposit      string
	Observations string
}
