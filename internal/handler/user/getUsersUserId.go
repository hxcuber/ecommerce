package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/hxcuber/ecommerce/internal/handler/user/response"
	"github.com/hxcuber/ecommerce/pkg/util"
	"net/http"
)

func (h Handler) GetUsersUserId() http.HandlerFunc {
	return util.ErrorHandler(func(w http.ResponseWriter, r *http.Request) (render.Renderer, int) {
		id := chi.URLParam(r, "user_id")
		if id == "" {
			return util.ErrorDesc{Description: "no id provided"}, http.StatusBadRequest
		}

		if !util.ValidUUID(id) {
			return util.ErrorDesc{Description: "unable to parse id"}, http.StatusBadRequest
		}

		user, err := h.ctrl.GetUserDetails(r.Context(), util.UUIDString(id))
		if err != nil {
			return util.ErrorDesc{Description: err.Error()}, http.StatusInternalServerError
		}

		return response.UserDetailsResponse{
			UserID:    string(user.UserID),
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		}, http.StatusOK

	})
}
