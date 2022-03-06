package attendanceHandler

import "net/http"

type AttHandlerInterface interface {
	GetAttendancesRangeDate(w http.ResponseWriter, r *http.Request)
	GetAttendancesCurrentUser(w http.ResponseWriter, r *http.Request)
	CreateAttendance(w http.ResponseWriter, r *http.Request)
	UpdateAttendance(w http.ResponseWriter, r *http.Request)
}
