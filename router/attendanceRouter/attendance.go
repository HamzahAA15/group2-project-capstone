package attendanceRouter

import (
	"net/http"
	"sirclo/project-capstone/controller/handler/attendanceHandler"
	"sirclo/project-capstone/controller/service/attendanceService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/repository/attendanceRepository"
	"sirclo/project-capstone/repository/userRepository"

	"github.com/gorilla/mux"
)

type AttResource struct{}

func (ar AttResource) AttRoute(attRepo attendanceRepository.AttendanceRepoInterface, userRepo userRepository.UserRepoInterface) *mux.Router {
	attService := attendanceService.NewAttendanceService(attRepo, userRepo)
	userService := userService.NewUserService(userRepo)
	attHandler := attendanceHandler.NewAttendanceHandler(attService, userService)

	router := mux.NewRouter()
	router.Handle("/", middleware.Authentication(http.HandlerFunc(attHandler.GetAttendances))).Methods("GET")
	router.Handle("/", middleware.Authentication(http.HandlerFunc(attHandler.GetAttendancesRangeDate))).Methods("GET")
	router.Handle("/", middleware.Authentication(http.HandlerFunc(attHandler.CreateAttendance))).Methods("POST")
	router.Handle("/", middleware.Authentication(http.HandlerFunc(attHandler.UpdateAttendance))).Methods("PUT")
	return router
}
