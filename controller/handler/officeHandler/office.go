package officeHandler

import (
	"encoding/json"
	"net/http"
	"sirclo/project-capstone/controller/service/officeService"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/response/officeResponse"
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
		response, _ := json.Marshal(utils.APIResponse("Something Went Wrong", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		var data []officeResponse.OfficeResponse
		for i := 0; i < len(offices); i++ {
			formatter := officeResponse.FormatOffice(offices[i])
			data = append(data, formatter)
		}

		response, _ := json.Marshal(utils.APIResponse("Success Get Offices Data", http.StatusOK, true, data))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
