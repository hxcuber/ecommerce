package product

import (
	"github.com/hxcuber/ecommerce/internal/controller/product"
)

type Handler struct {
	ctrl product.Controller
}

func NewHandler(ctrl product.Controller) Handler {
	return Handler{
		ctrl: ctrl,
	}
}
