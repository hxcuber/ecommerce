package response

import (
	"net/http"
	"time"
)

type GetProductsProductIdResponse struct {
	ProductDetailsResponse
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (resp GetProductsProductIdResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
