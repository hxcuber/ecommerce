package product

import (
	"github.com/hxcuber/ecommerce/pkg/db/pg"
)

type Repository interface {
	//GetProductByID(ctx context.Context, id uuid.UUID) (model.Product, error)
	//CreateProduct(ctx context.Context, product model.Product) (model.Product, error)
}

type impl struct {
	dbConn pg.ContextExecutor
}

func NewRepo(dbConn pg.ContextExecutor) Repository {
	return impl{
		dbConn: dbConn,
	}
}
