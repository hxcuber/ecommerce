package product

import (
	"context"
	"github.com/google/uuid"
	"github.com/hxcuber/ecommerce/internal/handler/product/request"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/hxcuber/ecommerce/pkg/util"
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
	},
}
