package logcatRouter

import (
	"net/http"
	"sirclo/project-capstone/controller/handler/logcatHandler"
	"sirclo/project-capstone/controller/service/logcatService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/repository/logcatRepository"
	"sirclo/project-capstone/repository/userRepository"

	"github.com/gorilla/mux"
)

type LogcatResource struct{}

func (lr LogcatResource) LogcatRoute(userRepo userRepository.UserRepoInterface, logcatRepo logcatRepository.LogcatRepoInterface) *mux.Router {
	userService := userService.NewUserService(userRepo)
	logcatService := logcatService.NewLogcatService(logcatRepo)
	logcatHandler := logcatHandler.NewLogcatHandler(userService, logcatService)

	router := mux.NewRouter()
	router.Handle("/", middleware.Authentication(http.HandlerFunc(logcatHandler.LogcatsHandler))).Methods("GET")
	router.Handle("/user", middleware.Authentication(http.HandlerFunc(logcatHandler.LogcatUserHandler))).Methods("GET")
	return router
}
