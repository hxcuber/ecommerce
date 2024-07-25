package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/hxcuber/ecommerce/internal/handler/user/response"
	"github.com/hxcuber/ecommerce/pkg/util"
	"net/http"
)

func (h Handler) GetUserDetails() http.HandlerFunc {
	return util.ErrorHandler(func(w http.ResponseWriter, r *http.Request) (render.Renderer, int) {
		idString := chi.URLParam(r, "user_id")
		if idString == "" {
			return util.ErrorDesc{Description: "no id provided"}, http.StatusBadRequest
		}

		id, err := uuid.Parse(idString)
		if err != nil {
			return util.ErrorDesc{Description: "unable to parse id"}, http.StatusBadRequest
		}

		user, err := h.ctrl.GetUserDetails(r.Context(), id)
		if err != nil {
			return util.ErrorDesc{Description: err.Error()}, http.StatusInternalServerError
		}

		return response.UserDetailsResponse{
			UserID:    user.UserID.String(),
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		}, http.StatusOK

	})
}
