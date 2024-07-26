package product

import (
	"github.com/go-chi/render"
	"github.com/hxcuber/ecommerce/internal/handler/product/request"
	"github.com/hxcuber/ecommerce/internal/handler/product/response"
	"github.com/hxcuber/ecommerce/pkg/util"
	"net/http"
)

func (h Handler) PostProducts() http.HandlerFunc {
	return util.ErrorHandler(func(w http.ResponseWriter, r *http.Request) (render.Renderer, int) {
		var req request.PostProductsRequest
		if err := render.Bind(r, &req); err != nil {
			return util.ErrorDesc{Description: err.Error()}, http.StatusBadRequest
		}

		product, err := h.ctrl.RegisterProduct(r.Context(), req)
		if err != nil {
			return util.ErrorDesc{Description: err.Error()}, http.StatusInternalServerError
		}

		return response.PostProductsResponse{
			ProductDetailsResponse: response.ProductDetailsResponse{
				ProductID:   string(product.ProductID),
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Stock:       product.Stock,
				CategoryID:  product.CategoryID,
			},
			CreatedAt: product.CreatedAt,
		}, http.StatusCreated
	})
}
