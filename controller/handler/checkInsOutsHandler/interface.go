package checkInsOutsHandler

import "net/http"

type CheckInOutHandlerInterface interface {
	GetsHandler(w http.ResponseWriter, r *http.Request)
	CheckinsHandler(w http.ResponseWriter, r *http.Request)
	CheckoutsHandler(w http.ResponseWriter, r *http.Request)
}
