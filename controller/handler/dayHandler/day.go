package dayHandler

import (
	"encoding/json"
	"net/http"
	"sirclo/project-capstone/controller/service/dayService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/request/dayRequest"
	"sirclo/project-capstone/utils/response/dayResponse"
	"sirclo/project-capstone/utils/validation"
)

type dayHandler struct {
	dayService  dayService.DayServiceInterface
	userService userService.UserServiceInterface
}

func NewDayHandler(dayService dayService.DayServiceInterface, userService userService.UserServiceInterface) DayHandlerInterface {
	return &dayHandler{
		dayService:  dayService,
		userService: userService,
	}
}

func (dh *dayHandler) GetDaysHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	officeID := queryParams.Get("office_id")
	date := queryParams.Get("date")

	days, err := dh.dayService.GetDays(officeID, date)
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
	ctx := r.Context()
	userID := middleware.ForContext(ctx).ID

	user, _ := dh.userService.GetUser(userID)
	if user.Role != "admin" {
		response, _ := json.Marshal(utils.APIResponse("You are not admin", http.StatusUnauthorized, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
		return
	}

	var input dayRequest.DayUpdateRequest
	json.NewDecoder(r.Body).Decode(&input)

	errValidation := validation.CheckEmpty(input.ID, input.Quota)
	if errValidation != nil {
		response, _ := json.Marshal(utils.APIResponse(errValidation.Error(), http.StatusUnprocessableEntity, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(response)
		return
	}

	dayUpdate, err := dh.dayService.UpdateDays(input)
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
