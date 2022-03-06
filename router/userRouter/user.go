package userRouter

import (
	"net/http"
	"sirclo/project-capstone/controller/handler/userHandler"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/repository/userRepository"

	"github.com/gorilla/mux"
)

type UserResource struct{}

func (ur UserResource) UserRoute(userRepo userRepository.UserRepoInterface) *mux.Router {
	userService := userService.NewUserService(userRepo)
	userHandler := userHandler.NewUserHandler(userService)

	router := mux.NewRouter()
	router.HandleFunc("/login", userHandler.Login).Methods("POST")
<<<<<<< HEAD
	router.HandleFunc("/register", userHandler.CreateUserHandler).Methods("POST")
=======
	router.Handle("/", middleware.Authentication(http.HandlerFunc(userHandler.CreateUserHandler))).Methods("POST")
>>>>>>> 7b50bc7e692324b69db179a1087fc11f0dee3064
	router.Handle("/", middleware.Authentication(http.HandlerFunc(userHandler.UpdateUserHandler))).Methods("PUT")
	router.Handle("/profile", middleware.Authentication(http.HandlerFunc(userHandler.GetUserHandler))).Methods("GET")
	router.Handle("/avatar", middleware.Authentication(http.HandlerFunc(userHandler.UploadFileHandler))).Methods("PUT")
	return router
}
