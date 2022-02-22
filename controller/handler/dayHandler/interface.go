package dayHandler

import "net/http"

type DayHandlerInterface interface {
	GetDaysHandler(w http.ResponseWriter, r *http.Request)
	UpdateDaysHandler(w http.ResponseWriter, r *http.Request)
}
