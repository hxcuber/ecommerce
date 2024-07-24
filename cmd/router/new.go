package router

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/controller/user"
	userHandler "github.com/hxcuber/ecommerce/internal/handler/user"
)

func New(ctx context.Context, userCtrl user.Controller) Router {
	return Router{
		userRESTHandler: userHandler.NewHandler(userCtrl),
	}
}
