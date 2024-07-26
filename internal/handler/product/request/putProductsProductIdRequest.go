package request

import (
	"net/http"
)

type PutProductsProductIdRequest PostProductsRequest

func (req PutProductsProductIdRequest) Bind(r *http.Request) error {
	return ProductDetailsRequest(req).Bind(r)
}
