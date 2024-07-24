package model

type OrderItem struct {
	OrderID   string `boil:"order_id" json:"order_id" toml:"order_id" yaml:"order_id"`
	ProductID string `boil:"product_id" json:"product_id" toml:"product_id" yaml:"product_id"`
	Quantity  int64  `boil:"quantity" json:"quantity" toml:"quantity" yaml:"quantity"`
}
