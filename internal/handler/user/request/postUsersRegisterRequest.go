package request

import (
	"github.com/friendsofgo/errors"
	"net/http"
)

type PostUsersRegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req PostUsersRegisterRequest) Bind(r *http.Request) error {
	if req.Username == "" {
		return errors.New("username is required")
	}
	if req.Email == "" {
		return errors.New("email is required")
	}
	if req.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
