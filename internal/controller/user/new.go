package user

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/handler/user/request"
	"github.com/hxcuber/ecommerce/internal/repository"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/hxcuber/ecommerce/pkg/util"
)

type Controller interface {
	GetUserDetails(ctx context.Context, id util.UUIDString) (model.User, error)
	RegisterUser(ctx context.Context, request request.PostUsersRegisterRequest) (model.User, error)
}

type impl struct {
	reg repository.Registry
}

func NewController(reg repository.Registry) Controller {
	return impl{
		reg: reg,
	}
}
