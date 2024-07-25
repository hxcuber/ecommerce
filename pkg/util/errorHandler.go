package util

import (
	"github.com/go-chi/render"
	"net/http"
)

func ErrorHandler(f func(w http.ResponseWriter, r *http.Request) (render.Renderer, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, stat := f(w, r)
		render.Status(r, stat)
		render.Render(w, r, resp)
	}
}
