package params

type CreateProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Sku         string `json:"sku"`
	Stock       int64  `json:"stock"`
	Price       int64  `json:"price"`
}
