package dayHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sirclo/project-capstone/controller/service/dayService"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/request/dayRequest"
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

func (dh *dayHandler) UpdateDaysHandler(w http.ResponseWriter, r *http.Request) {
	var input dayRequest.DayUpdateRequest
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&input)

	dayUpdate, err := dh.dayService.UpdateDays(input)
	fmt.Println("DU & err: ", dayUpdate, err)
	switch {
	case err != nil:
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		formatter := dayResponse.FormatUpdateDay(dayUpdate)
		response, _ := json.Marshal(utils.APIResponse("Success Update Day Data", http.StatusOK, true, formatter))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

}
