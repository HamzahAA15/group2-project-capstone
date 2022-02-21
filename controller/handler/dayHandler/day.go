package dayHandler

import (
	"encoding/json"
	"net/http"
	"sirclo/project-capstone/controller/service/dayService"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/response/dayResponse"
)

type dayHandler struct {
	dayService dayService.DayServiceInterface
}

func NewDayHandler(dayService dayService.DayServiceInterface) DayHandlerInterface {
	return &dayHandler{
		dayService: dayService,
	}
}

func (dh *dayHandler) GetDaysHandler(w http.ResponseWriter, r *http.Request) {
	days, err := dh.dayService.GetDays()
	switch {
	case err != nil:
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		var data []dayResponse.DayResponse
		for _, val := range days {
			dayFormatter := dayResponse.FormatDay(val)
			data = append(data, dayFormatter)
		}

		response, _ := json.Marshal(utils.APIResponse("Success Get Days Data", http.StatusOK, true, data))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

}
