package product

import (
	"context"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i impl) UpdateProductById(ctx context.Context, product model.Product) (model.Product, error) {
	ormProduct := product.ToOrmProduct()
	_, err := ormProduct.Update(ctx, i.dbConn, boil.Blacklist())
	if err != nil {
		return model.Product{}, err
	}

	return model.FromOrmProduct(ormProduct)
}
