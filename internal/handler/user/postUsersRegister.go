package user

import (
	"errors"
	"github.com/go-chi/render"
	userCtrl "github.com/hxcuber/ecommerce/internal/controller/user"
	"github.com/hxcuber/ecommerce/internal/handler/user/request"
	"github.com/hxcuber/ecommerce/internal/handler/user/response"
	"github.com/hxcuber/ecommerce/pkg/util"
	"net/http"
)

func (h Handler) PostUsersRegister() http.HandlerFunc {
	return util.ErrorHandler(func(w http.ResponseWriter, r *http.Request) (render.Renderer, int) {
		var req request.RegisterUserRequest
		if err := render.Bind(r, &req); err != nil {
			return util.ErrorDesc{Description: err.Error()}, http.StatusBadRequest
		}

		user, err := h.ctrl.RegisterUser(r.Context(), req)
		if errors.Is(err, userCtrl.ErrEmailInUse) || errors.Is(err, userCtrl.ErrUsernameInUse) {
			return util.ErrorDesc{Description: err.Error()}, http.StatusBadRequest
		}
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
