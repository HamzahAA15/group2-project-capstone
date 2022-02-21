package dayRouter

import (
	"sirclo/project-capstone/controller/handler/dayHandler"
	"sirclo/project-capstone/controller/service/dayService"
	"sirclo/project-capstone/repository/dayRepository"

	"github.com/gorilla/mux"
)

type DayResource struct{}

func (dr DayResource) DayRoute(dayRepo dayRepository.DayRepoInterface) *mux.Router {
	dayService := dayService.NewDayService(dayRepo)
	dayHandler := dayHandler.NewDayHandler(dayService)

	router := mux.NewRouter()
	router.HandleFunc("/", dayHandler.GetDaysHandler).Methods("GET")
	return router
}
