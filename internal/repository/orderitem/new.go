package orderitem

import (
	"github.com/hxcuber/ecommerce/pkg/db/pg"
)

type Repository interface {
}

type impl struct {
	dbConn pg.ContextExecutor
}

func NewRepo(dbConn pg.ContextExecutor) Repository {
	return impl{
		dbConn: dbConn,
	}
}
