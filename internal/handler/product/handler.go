package product

import (
	"github.com/hxcuber/ecommerce/internal/controller/product"
)

type Handler struct {
	productCtrl product.Controller
}

func NewHandler(productCtrl product.Controller) Handler {
	return Handler{
		productCtrl: productCtrl,
	}
}
