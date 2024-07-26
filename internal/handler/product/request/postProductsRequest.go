package request

import (
	"net/http"
)

type PostProductsRequest ProductDetailsRequest

func (req PostProductsRequest) Bind(r *http.Request) error {
	return ProductDetailsRequest(req).Bind(r)
}
