package response

import "net/http"

type PostUsersRegisterResponse UserDetailsResponse

func (resp PostUsersRegisterResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
