package router

import (
	"net/http"
	"sirclo/project-capstone/repository/userRepository"
	"sirclo/project-capstone/router/userRouter"
	"strings"

	"github.com/gorilla/mux"
)

func Routes(
	userRepo userRepository.UserRepoInterface,
) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	mount(router, "/users", userRouter.UserResource{}.UserRoute(userRepo))

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
