package attendanceHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sirclo/project-capstone/controller/service/attendanceService"
	"sirclo/project-capstone/controller/service/logcatService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/request/attendanceRequest"
	"sirclo/project-capstone/utils/response/attendanceResponse"
	"sirclo/project-capstone/utils/validation"
)

type attHandler struct {
	attService    attendanceService.AttServiceInterface
	userService   userService.UserServiceInterface
	logcatService logcatService.LogcatServiceInterface
}

func NewAttendanceHandler(attService attendanceService.AttServiceInterface, userService userService.UserServiceInterface, logcatService logcatService.LogcatServiceInterface) AttHandlerInterface {
	return &attHandler{
		attService:    attService,
		userService:   userService,
		logcatService: logcatService,
	}
}

func (ah *attHandler) GetAttendancesRangeDate(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	employeeEmail := queryParams.Get("employee_email")
	dateStart := queryParams.Get("date_start")
	dateEnd := queryParams.Get("date_end")
	status := queryParams.Get("status")
	officeId := queryParams.Get("office_id")
	order := queryParams.Get("order_by")

	attendances, err := ah.attService.GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, officeId, order)
	switch {
	case err != nil:
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		var data []attendanceResponse.AttGetResponse
		for _, val := range attendances {
			attFormatter := attendanceResponse.FormatGetAtt(val)
			data = append(data, attFormatter)
		}

		response, _ := json.Marshal(utils.APIResponse("Success Get Attendances Data", http.StatusOK, true, data))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (ah *attHandler) GetAttendancesCurrentUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	status := queryParams.Get("status")
	order := queryParams.Get("order_by")
	if order == "" {
		order = "desc"
	}
	ctx := r.Context()
	user := middleware.ForContext(ctx)

	attendances, err := ah.attService.GetAttendancesCurrentUser(user.ID, status, order)
	switch {
	case err != nil:
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		var count int
		var dataMain attendanceResponse.AttGetUserResp
		var data []attendanceResponse.AttGetResponse
		for _, val := range attendances {
			attFormatter := attendanceResponse.FormatGetAtt(val)
			data = append(data, attFormatter)
			count++
		}
		dataMain.AttGetResponse = data
		dataMain.Total = count

		response, _ := json.Marshal(utils.APIResponse("Success Get Attendances Data", http.StatusOK, true, dataMain))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (ah *attHandler) CreateAttendance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForContext(ctx)

	var input attendanceRequest.CreateAttRequest
	json.NewDecoder(r.Body).Decode(&input)

	errValidation := validation.CheckEmpty(input.Day)
	if errValidation != nil {
		response, _ := json.Marshal(utils.APIResponse(errValidation.Error(), http.StatusUnprocessableEntity, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(response)
		return
	}

	_, err := ah.attService.CreateAttendance(user.ID, input)
	switch {
	case err != nil: // error internal server
		response, _ := json.Marshal(utils.APIResponse(err.Error(), http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		GetUser, _ := ah.userService.GetUser(user.ID)
		message := fmt.Sprintf("%s have requested for WFO", GetUser.Name)
		ah.logcatService.CreateLogcat(user.ID, message, "attendances")
		response, _ := json.Marshal(utils.APIResponse("Success Create Attendace Data", http.StatusOK, true, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

}

func (ah *attHandler) UpdateAttendance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForContext(ctx)

	CurrentUser, err := ah.userService.GetUser(user.ID)
	if err != nil {
		response, _ := json.Marshal(utils.APIResponse(err.Error(), http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	}
	if CurrentUser.Role != "admin" {
		response, _ := json.Marshal(utils.APIResponse("You are not admin", http.StatusUnauthorized, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
		return
	}

	var input attendanceRequest.UpdateAttRequest
	json.NewDecoder(r.Body).Decode(&input)

	errValidation := validation.CheckEmpty(input.ID, input.Status)
	if errValidation != nil {
		response, _ := json.Marshal(utils.APIResponse(errValidation.Error(), http.StatusUnprocessableEntity, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(response)
		return
	}

	_, errUpdate := ah.attService.UpdateAttendance(user.ID, input)
	switch {
	case errUpdate != nil:
		response, _ := json.Marshal(utils.APIResponse(err.Error(), http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		userId, employeeName, _ := ah.attService.GetAttendancesById(input.ID)
		message := fmt.Sprintf("%s have updated request status on %s", CurrentUser.Name, employeeName)
		ah.logcatService.CreateLogcat(userId, message, "attendances")
		response, _ := json.Marshal(utils.APIResponse("Success Update Day Data", http.StatusOK, true, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
