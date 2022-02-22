package certificateRouter

import (
	"net/http"
	"sirclo/project-capstone/controller/handler/certificateHandler"
	"sirclo/project-capstone/controller/service/certificateService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/repository/certificateRepository"
	"sirclo/project-capstone/repository/userRepository"

	"github.com/gorilla/mux"
)

type CertificateResource struct{}

func (cr CertificateResource) CertificateRoute(
	certificateRepo certificateRepository.CertificateInterface,
	userRepo userRepository.UserRepoInterface,
) *mux.Router {
	certificateService := certificateService.NewCertificateService(certificateRepo)
	userService := userService.NewUserService(userRepo)

	certificateHandler := certificateHandler.NewCertificateHandler(certificateService, userService)

	router := mux.NewRouter()
	router.Handle("/", middleware.Authentication(http.HandlerFunc(certificateHandler.GetCertificatesHandler))).Methods("GET")
	return router
}
