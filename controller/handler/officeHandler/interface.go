package officeHandler

import "net/http"

type OfficeHandlerInterface interface {
	GetOfficesHandler(w http.ResponseWriter, r *http.Request)
}
