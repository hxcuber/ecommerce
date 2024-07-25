package router

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/hxcuber/ecommerce/internal/handler/user"
	"net/http"
)

type Router struct {
	ctx             context.Context
	userRESTHandler user.Handler
}

func (rtr Router) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,       // log relationship request calls
		middleware.StripSlashes, // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,    // recover from panics without crashing server
		middleware.RequestID,    // unique request ids for logging
		middleware.RealIP,       // differentiate hosts
	)

	r.Route("/users", func(r chi.Router) {
		r.Get("/{user_id}", rtr.userRESTHandler.GetUsersUserId())
		r.Post("/register", rtr.userRESTHandler.PostUsersRegister())
	})

	return r
}
