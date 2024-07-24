package user

import "github.com/hxcuber/ecommerce/internal/controller/user"

type Handler struct {
	ctrl user.Controller
}

func NewHandler(ctrl user.Controller) Handler {
	return Handler{ctrl: ctrl}
}
