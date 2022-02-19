package officeRouter

import (
	"sirclo/project-capstone/controller/handler/officeHandler"
	"sirclo/project-capstone/controller/service/officeService"
	"sirclo/project-capstone/repository/officeRepository"

	"github.com/gorilla/mux"
)

type OfficeResource struct{}

func (ur OfficeResource) OfficeRoute(officeRepo officeRepository.OfficeRepoInterface) *mux.Router {
	officeService := officeService.NewOfficeService(officeRepo)
	officeHandler := officeHandler.NewOfficeHandler(officeService)

	router := mux.NewRouter()
	router.HandleFunc("/", officeHandler.GetOfficesHandler).Methods("GET")
	return router
}
