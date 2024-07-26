package response

type ProductDetailsResponse struct {
	ProductID   string  `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       int64   `json:"stock"`
	CategoryID  int64   `json:"category_id"`
}
