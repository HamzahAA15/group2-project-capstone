package attendanceRouter

import (
	"net/http"
	"sirclo/project-capstone/controller/handler/attendanceHandler"
	"sirclo/project-capstone/controller/service/attendanceService"
	"sirclo/project-capstone/controller/service/logcatService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/repository/attendanceRepository"
	"sirclo/project-capstone/repository/logcatRepository"
	"sirclo/project-capstone/repository/userRepository"

	"github.com/gorilla/mux"
)

type AttResource struct{}

func (ar AttResource) AttRoute(attRepo attendanceRepository.AttendanceRepoInterface, userRepo userRepository.UserRepoInterface, logcatRepo logcatRepository.LogcatRepoInterface) *mux.Router {
	attService := attendanceService.NewAttendanceService(attRepo, userRepo)
	userService := userService.NewUserService(userRepo)
	logcatService := logcatService.NewLogcatService(logcatRepo)
	attHandler := attendanceHandler.NewAttendanceHandler(attService, userService, logcatService)

	router := mux.NewRouter()
	router.Handle("/", middleware.Authentication(http.HandlerFunc(attHandler.GetAttendancesRangeDate))).Methods("GET")
	router.Handle("/user", middleware.Authentication(http.HandlerFunc(attHandler.GetAttendancesCurrentUser))).Methods("GET")
	router.Handle("/ischeckin", middleware.Authentication(http.HandlerFunc(attHandler.IsCheckins))).Methods("GET")
	router.Handle("/", middleware.Authentication(http.HandlerFunc(attHandler.CreateAttendance))).Methods("POST")
	router.Handle("/", middleware.Authentication(http.HandlerFunc(attHandler.UpdateAttendance))).Methods("PUT")
	return router
}
