package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/hxcuber/ecommerce/internal/repository"
	"github.com/hxcuber/ecommerce/pkg/model"
)

type Controller interface {
	GetUserDetails(ctx context.Context, id uuid.UUID) (model.User, error)
	RegisterUser(ctx context.Context, user model.User, password string) (model.User, error)
}

type impl struct {
	reg repository.Registry
}

func NewController(reg repository.Registry) Controller {
	return impl{
		reg: reg,
	}
}
