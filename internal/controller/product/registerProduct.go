package product

import (
	"context"
	"github.com/google/uuid"
	"github.com/hxcuber/ecommerce/internal/handler/product/request"
	"github.com/hxcuber/ecommerce/internal/repository"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"github.com/hxcuber/ecommerce/pkg/logerr"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/hxcuber/ecommerce/pkg/util"
	"log"
	"time"
)

func (i impl) RegisterProduct(ctx context.Context, request request.PostProductsRequest) (model.Product, error) {
	product := model.Product{
		ProductID:   util.UUIDString(uuid.New().String()),
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		CategoryID:  request.CategoryID,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	var returnProduct *orm.Product
	err := i.reg.DoInTx(context.Background(), func(ctx context.Context, registry repository.Registry) error {
		var err error
		returnProduct, err = registry.Product().CreateProduct(ctx, product)
		if err != nil {
			log.Printf(logerr.LogErrMessage("RegisterProduct", "creating product", err))
			return err
		}
		return nil
	}, nil)

	if err != nil {
		log.Printf(logerr.LogErrMessage("RegisterProduct", "doing in transaction", err))
		return model.Product{}, err
	}

	return model.FromOrmProduct(returnProduct), nil
}
