package response

import (
	"net/http"
	"time"
)

type UserDetailsResponse struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (resp UserDetailsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
