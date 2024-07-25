package product

import (
	"context"
	"github.com/google/uuid"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"github.com/hxcuber/ecommerce/pkg/model"
)

func (i impl) GetProductByID(ctx context.Context, id uuid.UUID) (model.Product, error) {
	ormProduct, err := orm.FindProduct(ctx, i.dbConn, id.String())
	if err != nil {
		return model.Product{}, err
	}

	return model.FromOrmProduct(ormProduct)
}
