package types

type MovementWithProductId struct {
	Movement  Movement `json:"movement"`
	ProductId int64    `json:"product_id"`
}
