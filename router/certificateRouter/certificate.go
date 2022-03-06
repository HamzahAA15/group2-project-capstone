package certificateRouter

import (
	"net/http"
	"sirclo/project-capstone/controller/handler/certificateHandler"
	"sirclo/project-capstone/controller/service/certificateService"
	"sirclo/project-capstone/controller/service/logcatService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/repository/certificateRepository"
	"sirclo/project-capstone/repository/logcatRepository"
	"sirclo/project-capstone/repository/userRepository"

	"github.com/gorilla/mux"
)

type CertificateResource struct{}

func (cr CertificateResource) CertificateRoute(
	certificateRepo certificateRepository.CertificateInterface,
	userRepo userRepository.UserRepoInterface,
	logcatRepo logcatRepository.LogcatRepoInterface,
) *mux.Router {
	certificateService := certificateService.NewCertificateService(certificateRepo)
	userService := userService.NewUserService(userRepo)
	logcatService := logcatService.NewLogcatService(logcatRepo)

	certificateHandler := certificateHandler.NewCertificateHandler(certificateService, userService, logcatService)

	router := mux.NewRouter()
	router.Handle("/", middleware.Authentication(http.HandlerFunc(certificateHandler.GetCertificatesHandler))).Methods("GET")
	router.Handle("/user", middleware.Authentication(http.HandlerFunc(certificateHandler.GetCertificateUserHandler))).Methods("GET")
	router.Handle("/", middleware.Authentication(http.HandlerFunc(certificateHandler.UploadCertificateHandler))).Methods("POST")
	router.Handle("/", middleware.Authentication(http.HandlerFunc(certificateHandler.VerifyCertificateHandler))).Methods("PUT")
	return router
}
