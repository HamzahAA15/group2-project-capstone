package dayRouter

import (
	"net/http"
	"sirclo/project-capstone/controller/handler/dayHandler"
	"sirclo/project-capstone/controller/service/dayService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/repository/dayRepository"
	"sirclo/project-capstone/repository/userRepository"

	"github.com/gorilla/mux"
)

type DayResource struct{}

func (dr DayResource) DayRoute(dayRepo dayRepository.DayRepoInterface, userRepo userRepository.UserRepoInterface) *mux.Router {
	dayService := dayService.NewDayService(dayRepo, userRepo)
	userService := userService.NewUserService(userRepo)
	dayHandler := dayHandler.NewDayHandler(dayService, userService)

	router := mux.NewRouter()
	router.HandleFunc("/", dayHandler.GetDaysHandler).Methods("GET")
	router.Handle("/", middleware.Authentication(http.HandlerFunc(dayHandler.UpdateDaysHandler))).Methods("PUT")
	return router
}
