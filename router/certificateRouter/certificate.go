package certificateRouter

import (
	"sirclo/project-capstone/controller/handler/certificateHandler"
	"sirclo/project-capstone/controller/service/certificateService"
	"sirclo/project-capstone/repository/certificateRepository"

	"github.com/gorilla/mux"
)

type CertificateResource struct{}

func (cr CertificateResource) CertificateRoute(certificateRepo certificateRepository.CertificateInterface) *mux.Router {
	certificateService := certificateService.NewCertificateService(certificateRepo)
	certificateHandler := certificateHandler.NewCertificateHandler(certificateService)

	router := mux.NewRouter()
	router.HandleFunc("/", certificateHandler.GetCertificatesHandler).Methods("GET")
	return router
}
