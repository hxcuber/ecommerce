package product

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/handler/product/request"
	"github.com/hxcuber/ecommerce/internal/repository"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/hxcuber/ecommerce/pkg/util"
)

type Controller interface {
	GetProductDetails(ctx context.Context, id util.UUIDString) (model.Product, error)
	RegisterProduct(ctx context.Context, request request.PostProductsRequest) (model.Product, error)
	UpdateProduct(ctx context.Context, id util.UUIDString, request request.PutProductsProductIdRequest) (model.Product, error)
}

type impl struct {
	reg repository.Registry
}

func NewController(reg repository.Registry) Controller {
	return impl{
		reg: reg,
	}
}
