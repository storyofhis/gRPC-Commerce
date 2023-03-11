package params

type CreateOrder struct {
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}
