package userHandler

import "net/http"

type UserHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	GetUsersHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
	UpdateUserHandler(w http.ResponseWriter, r *http.Request)
	DeleteUserHandler(w http.ResponseWriter, r *http.Request)
	UploadFileHandler(w http.ResponseWriter, r *http.Request)
}
