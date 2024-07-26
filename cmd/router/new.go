package router

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/controller/product"
	"github.com/hxcuber/ecommerce/internal/controller/user"
	productHandler "github.com/hxcuber/ecommerce/internal/handler/product"
	userHandler "github.com/hxcuber/ecommerce/internal/handler/user"
)

func New(ctx context.Context, userCtrl user.Controller, productCtrl product.Controller) Router {
	return Router{
		ctx:                ctx,
		userRESTHandler:    userHandler.NewHandler(userCtrl),
		productRESTHandler: productHandler.NewHandler(productCtrl),
	}
}
