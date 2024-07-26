package product

import (
	"context"
	"github.com/hxcuber/ecommerce/pkg/logerr"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/hxcuber/ecommerce/pkg/util"
	"log"
)

func (i impl) GetProductDetails(ctx context.Context, id util.UUIDString) (model.Product, error) {
	product, err := i.reg.Product().GetProductByID(context.Background(), id)
	if err != nil {
		log.Printf(logerr.LogErrMessage("GetProductDetails", "getting product %s", err, string(id)))
		return model.Product{}, err
	}

	return model.FromOrmProduct(product), nil
}
