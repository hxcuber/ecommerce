package user

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"github.com/hxcuber/ecommerce/pkg/db/pg"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/hxcuber/ecommerce/pkg/util"
)

type Repository interface {
	GetUserById(ctx context.Context, id util.UUIDString) (*orm.User, error)
	GetUserByEmail(ctx context.Context, email string) (*orm.User, error)
	GetUserByUsername(ctx context.Context, username string) (*orm.User, error)
	CreateUser(ctx context.Context, user model.User) (*orm.User, error)
}

type impl struct {
	dbConn pg.ContextExecutor
}

func NewRepo(dbConn pg.ContextExecutor) Repository {
	return impl{
		dbConn: dbConn,
	}
}
