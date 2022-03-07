package checkInOutRouter

import (
	"net/http"
	"sirclo/project-capstone/controller/handler/checkInsOutsHandler"
	"sirclo/project-capstone/controller/service/checkInsOutsService"
	"sirclo/project-capstone/controller/service/logcatService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/repository/checkInOutRepository"
	"sirclo/project-capstone/repository/logcatRepository"
	"sirclo/project-capstone/repository/userRepository"

	"github.com/gorilla/mux"
)

type CheckInOutResource struct{}

func (cr CheckInOutResource) CheckInOutRoute(
	checkInOutRepo checkInOutRepository.CheckInOutRepoInterface,
	userRepo userRepository.UserRepoInterface,
	logcatRepo logcatRepository.LogcatRepoInterface,
) *mux.Router {
	checkInOutService := checkInsOutsService.NewCheckInOutService(checkInOutRepo)
	userService := userService.NewUserService(userRepo)
	logcatService := logcatService.NewLogcatService(logcatRepo)
	checkInOutHandler := checkInsOutsHandler.NewCheckInOutHandler(checkInOutService, userService, logcatService)

	router := mux.NewRouter()
	router.Handle("/", middleware.Authentication(http.HandlerFunc(checkInOutHandler.GetsHandler))).Methods("GET")
	router.Handle("/user", middleware.Authentication(http.HandlerFunc(checkInOutHandler.GetsByUserHandler))).Methods("GET")
	router.Handle("/ins", middleware.Authentication(http.HandlerFunc(checkInOutHandler.CheckinsHandler))).Methods("POST")
	router.Handle("/outs", middleware.Authentication(http.HandlerFunc(checkInOutHandler.CheckoutsHandler))).Methods("POST")
	return router
}
