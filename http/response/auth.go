package response

import "github.com/androzd/fingo/model"

type UserNotFoundOrPasswordIsWrong struct {
	ErrorResponse
}

type UserData struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}
type UserLoggedIn struct {
	SuccessResponse
	Data UserData `json:"data"`
}
