package checkInOutRouter

import (
	"net/http"
	"sirclo/project-capstone/controller/handler/checkInsOutsHandler"
	"sirclo/project-capstone/controller/service/checkInsOutsService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/repository/checkInOutRepository"

	"github.com/gorilla/mux"
)

type CheckInOutResource struct{}

func (cr CheckInOutResource) CheckInOutRoute(
	checkInOutRepo checkInOutRepository.CheckInOutRepoInterface,
) *mux.Router {
	checkInOutService := checkInsOutsService.NewCheckInOutService(checkInOutRepo)
	checkInOutHandler := checkInsOutsHandler.NewCheckInOutHandler(checkInOutService)

	router := mux.NewRouter()
	router.Handle("/", middleware.Authentication(http.HandlerFunc(checkInOutHandler.GetsHandler))).Methods("GET")
	return router
}
