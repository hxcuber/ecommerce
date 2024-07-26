package request

import (
	"errors"
	"net/http"
)

type ProductDetailsRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price"`
	Stock       int64   `json:"stock"`
	CategoryID  int64   `json:"category_id"`
}

func (req ProductDetailsRequest) Bind(r *http.Request) error {
	if req.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
