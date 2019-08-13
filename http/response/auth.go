package response

import "github.com/androzd/finance/model"

type UserNotFoundOrPasswordIsWrong struct {
	ErrorResponse
}

type UserData struct {
	User model.User `json:"user"`
	Token string `json:"token"`
}
type UserLoggedIn struct {
	SuccessResponse
	Data UserData `json:"data"`
}