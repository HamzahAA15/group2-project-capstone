package officeHandler

import (
	"encoding/json"
	"net/http"
	"sirclo/project-capstone/controller/service/officeService"
	"sirclo/project-capstone/utils"
)

type officeHandler struct {
	officeService officeService.OfficeServiceInterface
}

func NewOfficeHandler(officeService officeService.OfficeServiceInterface) OfficeHandlerInterface {
	return &officeHandler{
		officeService: officeService,
	}
}

func (uh *officeHandler) GetOfficesHandler(w http.ResponseWriter, r *http.Request) {
	offices, err := uh.officeService.GetOffices()
	switch {
	case err != nil: // error internal server
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		response, _ := json.Marshal(utils.APIResponse("Success Get Offices Data", http.StatusOK, true, offices))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
