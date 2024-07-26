package product

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"github.com/hxcuber/ecommerce/pkg/db/pg"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/hxcuber/ecommerce/pkg/util"
)

type Repository interface {
	GetProductByID(ctx context.Context, id util.UUIDString) (*orm.Product, error)
	CreateProduct(ctx context.Context, product model.Product) (*orm.Product, error)
	UpdateProductById(ctx context.Context, product model.Product) (*orm.Product, error)
}

type impl struct {
	dbConn pg.ContextExecutor
}

func NewRepo(dbConn pg.ContextExecutor) Repository {
	return impl{
		dbConn: dbConn,
	}
}
