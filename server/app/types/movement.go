package types

import "time"

type Movement struct {
	id           int
	date         time.Time
	shippingCode string
	pallets      int
	units        int
	deposit      string
	observations string
}
