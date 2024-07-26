package response

import (
	"net/http"
	"time"
)

type PutProductsProductIdResponse struct {
	ProductDetailsResponse
	UpdatedAt time.Time `json:"updated_at"`
}

func (resp PutProductsProductIdResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
