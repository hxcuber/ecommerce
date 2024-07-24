package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/hxcuber/ecommerce/pkg/db/pg"
	"github.com/hxcuber/ecommerce/pkg/model"
)

type Repository interface {
	GetUserById(ctx context.Context, id uuid.UUID) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}

type impl struct {
	dbConn pg.ContextExecutor
}

func New(dbConn pg.ContextExecutor) Repository {
	return impl{
		dbConn: dbConn,
	}
}
