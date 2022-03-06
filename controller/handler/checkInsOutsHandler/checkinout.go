package checkInsOutsHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sirclo/project-capstone/controller/service/checkInsOutsService"
	"sirclo/project-capstone/controller/service/logcatService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/request/checkInsOutsRequest"
	"sirclo/project-capstone/utils/response/checkinsoutsResponse"
	"sirclo/project-capstone/utils/validation"
)

type checkInOutHandler struct {
	checkInOutService checkInsOutsService.CheckinoutServiceInterface
	userService       userService.UserServiceInterface
	logcatService     logcatService.LogcatServiceInterface
}

func NewCheckInOutHandler(checkInOutService checkInsOutsService.CheckinoutServiceInterface, userService userService.UserServiceInterface, logcatService logcatService.LogcatServiceInterface) CheckInOutHandlerInterface {
	return &checkInOutHandler{
		checkInOutService: checkInOutService,
		userService:       userService,
		logcatService:     logcatService,
	}
}

func (ch *checkInOutHandler) GetsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := middleware.ForContext(ctx)
	if userID == nil {
		response, _ := json.Marshal(utils.APIResponse("Unautorized", http.StatusUnauthorized, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
	}

	CheckInsOuts, err := ch.checkInOutService.Gets()
	switch {
	case err != nil:
		response, _ := json.Marshal(utils.APIResponse("Something Went Wrong", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		var data []checkinsoutsResponse.CheckInsOutsResponse
		for _, val := range CheckInsOuts {
			dayFormatter := checkinsoutsResponse.FormatCheckInsOuts(val)
			data = append(data, dayFormatter)
		}

		response, _ := json.Marshal(utils.APIResponse("Success Get All Data", http.StatusOK, true, data))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (ch *checkInOutHandler) CheckinsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := middleware.ForContext(ctx).ID

	var input checkInsOutsRequest.CheckInsRequest
	json.NewDecoder(r.Body).Decode(&input)

	errValidation := validation.CheckEmpty(input.AttendanceID, input.Temprature)
	if errValidation != nil {
		response, _ := json.Marshal(utils.APIResponse(errValidation.Error(), http.StatusUnprocessableEntity, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(response)
		return
	}

	userRequest, _ := ch.checkInOutService.CheckRequest(input.AttendanceID)
	if userRequest.Attendance.Employee.ID != userID || userRequest.Attendance.Status != "approved" {
		response, _ := json.Marshal(utils.APIResponse("you don't have permission to check-ins in this presence", http.StatusForbidden, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write(response)
		return
	}

	user := ch.checkInOutService.CheckData(userID, input.AttendanceID)
	if user > 0 {
		response, _ := json.Marshal(utils.APIResponse("you have been check-ins in this presence", http.StatusBadRequest, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	_, err := ch.checkInOutService.Checkin(input)
	switch {
	case err != nil: // error internal server
		response, _ := json.Marshal(utils.APIResponse("Something Went Wrong", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		GetUser, _ := ch.userService.GetUser(userID)
		message := fmt.Sprintf("%s have check in", GetUser.Name)
		ch.logcatService.CreateLogcat(userID, message, "checkin")
		response, _ := json.Marshal(utils.APIResponse("Check-Ins Success", http.StatusOK, true, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (ch *checkInOutHandler) CheckoutsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := middleware.ForContext(ctx).ID

	var input checkInsOutsRequest.CheckOutsRequest
	json.NewDecoder(r.Body).Decode(&input)

	errValidation := validation.CheckEmpty(input.ID, input.AttendanceID)
	if errValidation != nil {
		response, _ := json.Marshal(utils.APIResponse(errValidation.Error(), http.StatusUnprocessableEntity, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(response)
		return
	}

	_, err := ch.checkInOutService.Checkout(userID, input)
	switch {
	case err != nil: // error internal server
		response, _ := json.Marshal(utils.APIResponse("Something Went Wrong", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		GetUser, _ := ch.userService.GetUser(userID)
		message := fmt.Sprintf("%s have check out", GetUser.Name)
		ch.logcatService.CreateLogcat(userID, message, "checkout")
		response, _ := json.Marshal(utils.APIResponse("Check-Outs Success", http.StatusOK, true, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
