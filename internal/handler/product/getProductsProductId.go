package product

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/hxcuber/ecommerce/internal/handler/product/response"
	"github.com/hxcuber/ecommerce/pkg/util"
	"net/http"
)

func (h Handler) GetProductsProductId() http.HandlerFunc {
	return util.ErrorHandler(func(w http.ResponseWriter, r *http.Request) (render.Renderer, int) {
		id := chi.URLParam(r, "product_id")
		if id == "" {
			return util.ErrorDesc{Description: "no id provided"}, http.StatusBadRequest
		}
		if !util.ValidUUID(id) {
			return util.ErrorDesc{Description: "id not valid"}, http.StatusBadRequest
		}

		product, err := h.ctrl.GetProductDetails(r.Context(), util.UUIDString(id))
		if err != nil {
			return util.ErrorDesc{Description: err.Error()}, http.StatusInternalServerError
		}

		return response.GetProductsProductIdResponse{
			ProductDetailsResponse: response.ProductDetailsResponse{
				ProductID:   string(product.ProductID),
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Stock:       product.Stock,
				CategoryID:  product.CategoryID,
			},
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		}, http.StatusOK
	})
}
