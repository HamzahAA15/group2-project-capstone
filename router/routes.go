package router

import (
	"net/http"
	"sirclo/project-capstone/repository/certificateRepository"
	"sirclo/project-capstone/repository/dayRepository"
	"sirclo/project-capstone/repository/officeRepository"
	"sirclo/project-capstone/repository/userRepository"
	"sirclo/project-capstone/router/certificateRouter"
	"sirclo/project-capstone/router/dayRouter"
	"sirclo/project-capstone/router/officeRouter"
	"sirclo/project-capstone/router/userRouter"
	"strings"

	"github.com/gorilla/mux"
)

func Routes(
	userRepo userRepository.UserRepoInterface,
	officeRepo officeRepository.OfficeRepoInterface,
	certificateRepo certificateRepository.CertificateInterface,
	dayRepo dayRepository.DayRepoInterface,
) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	mount(router, "/users", userRouter.UserResource{}.UserRoute(userRepo))
	mount(router, "/offices", officeRouter.OfficeResource{}.OfficeRoute(officeRepo))
	mount(router, "/days", dayRouter.DayResource{}.DayRoute(dayRepo))
	mount(router, "/certificates", certificateRouter.CertificateResource{}.CertificateRoute(certificateRepo, userRepo))

	return router
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
