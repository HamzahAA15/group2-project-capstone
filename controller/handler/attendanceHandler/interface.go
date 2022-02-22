package attendanceHandler

import "net/http"

type AttHandlerInterface interface {
	CreateAttendance(w http.ResponseWriter, r *http.Request)
}
