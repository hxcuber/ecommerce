package response

import (
	"net/http"
	"time"
)

type PostProductsResponse struct {
	ProductDetailsResponse
	CreatedAt time.Time `json:"created_at"`
}

func (resp PostProductsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
