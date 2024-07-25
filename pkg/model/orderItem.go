package model

import "github.com/hxcuber/ecommerce/pkg/util"

type OrderItem struct {
	OrderID   util.UUIDString `boil:"order_id" json:"order_id" toml:"order_id" yaml:"order_id"`
	ProductID util.UUIDString `boil:"product_id" json:"product_id" toml:"product_id" yaml:"product_id"`
	Quantity  int64           `boil:"quantity" json:"quantity" toml:"quantity" yaml:"quantity"`
}
