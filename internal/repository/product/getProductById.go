package product

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"github.com/hxcuber/ecommerce/pkg/util"
)

func (i impl) GetProductByID(ctx context.Context, id util.UUIDString) (*orm.Product, error) {
	ormProduct, err := orm.FindProduct(ctx, i.dbConn, string(id))
	if err != nil {
		return nil, err
	}

	return ormProduct, nil
}
