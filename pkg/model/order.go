package model

import (
	"github.com/hxcuber/ecommerce/pkg/util"
	"time"
)

type Order struct {
	OrderID         util.UUIDString `boil:"order_id" json:"order_id" toml:"order_id" yaml:"order_id"`
	UserID          util.UUIDString `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	TotalAmount     int64           `boil:"total_amount" json:"total_amount" toml:"total_amount" yaml:"total_amount"`
	ShippingAddress string          `boil:"shipping_address" json:"shipping_address" toml:"shipping_address" yaml:"shipping_address"`
	Status          Status          `boil:"status" json:"status" toml:"status" yaml:"status"`
	CreatedAt       time.Time       `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt       time.Time       `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
}
