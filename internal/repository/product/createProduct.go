package product

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i impl) CreateProduct(ctx context.Context, product model.Product) (*orm.Product, error) {
	ormProduct := product.ToOrmProduct()
	err := ormProduct.Insert(ctx, i.dbConn, boil.Blacklist())
	if err != nil {
		return nil, err
	}

	return ormProduct, nil
}
