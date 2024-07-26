package product

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/handler/product/request"
	"github.com/hxcuber/ecommerce/internal/repository"
	"github.com/hxcuber/ecommerce/pkg/logerr"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/hxcuber/ecommerce/pkg/util"
	"log"
)

func (i impl) UpdateProduct(ctx context.Context, id util.UUIDString, request request.PutProductsProductIdRequest) (model.Product, error) {
	ormProduct, err := i.reg.Product().GetProductByID(context.Background(), id)
	if err != nil {
		log.Printf(logerr.LogErrMessage("UpdateProduct", "finding product by id", err))
		return model.Product{}, err
	}

	// TODO check for pending orders that affect this product

	product := model.FromOrmProduct(ormProduct)

	product.Name = request.Name
	product.Description = request.Description
	product.Price = request.Price
	product.Stock = request.Stock
	product.CategoryID = request.CategoryID

	err = i.reg.DoInTx(ctx, func(ctx context.Context, registry repository.Registry) error {
		var err error
		ormProduct, err = registry.Product().UpdateProductById(ctx, product)
		if err != nil {
			log.Printf(logerr.LogErrMessage("UpdateProduct", "updating product %s", err, string(product.ProductID)))
			return err
		}
		return nil
	}, nil)

	if err != nil {
		log.Printf(logerr.LogErrMessage("UpdateProduct", "doing in transaction", err))
	}

	return model.FromOrmProduct(ormProduct), nil
}
