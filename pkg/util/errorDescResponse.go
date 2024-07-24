package util

import "net/http"

type ErrorDesc struct {
	Description string `json:"description"`
}

func (e ErrorDesc) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
