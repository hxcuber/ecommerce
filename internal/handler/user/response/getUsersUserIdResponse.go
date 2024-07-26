package response

import "net/http"

type GetUsersUserIdResponse UserDetailsResponse

func (resp GetUsersUserIdResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
