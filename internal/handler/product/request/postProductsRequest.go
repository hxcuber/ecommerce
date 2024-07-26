package request

import (
	"errors"
	"net/http"
)

type PostProductsRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price"`
	Stock       int64   `json:"stock"`
	CategoryID  int64   `json:"category_id"`
}

func (req PostProductsRequest) Bind(r *http.Request) error {
	if req.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
