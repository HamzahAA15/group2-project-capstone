package userHandler

import "net/http"

type UserHandlerInterface interface {
	GetUserHandler(w http.ResponseWriter, r *http.Request)
}
