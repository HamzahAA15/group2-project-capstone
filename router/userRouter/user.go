package userRouter

import (
	"sirclo/project-capstone/controller/handler/userHandler"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/repository/userRepository"

	"github.com/gorilla/mux"
)

type UserResource struct{}

func (ur UserResource) UserRoute(userRepo userRepository.UserRepoInterface) *mux.Router {
	userService := userService.NewUserService(userRepo)
	userHandler := userHandler.NewUserHandler(userService)

	router := mux.NewRouter()
	router.HandleFunc("/{id}", userHandler.GetUserHandler).Methods("GET")
	return router
}
