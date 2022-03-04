package attendanceHandler

import (
	"encoding/json"
	"net/http"
	"sirclo/project-capstone/controller/service/attendanceService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/request/attendanceRequest"
	"sirclo/project-capstone/utils/response/attendanceResponse"
	"sirclo/project-capstone/utils/validation"
)

type attHandler struct {
	attService  attendanceService.AttServiceInterface
	userService userService.UserServiceInterface
}

func NewAttendanceHandler(attService attendanceService.AttServiceInterface, userService userService.UserServiceInterface) AttHandlerInterface {
	return &attHandler{
		attService:  attService,
		userService: userService,
	}
}

func (ah *attHandler) GetAttendances(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	employee := queryParams.Get("employee")
	date := queryParams.Get("date")
	status := queryParams.Get("status")
	office := queryParams.Get("office")
	order := queryParams.Get("order")
	if order == "" {
		order = "asc"
	}

	attendances, err := ah.attService.GetAttendances(employee, date, status, office, order)
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

func (ah *attHandler) GetAttendancesRangeDate(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	employeeEmail := queryParams.Get("employee_email")
	dateStart := queryParams.Get("date_start")
	dateEnd := queryParams.Get("date_end")
	status := queryParams.Get("status")
	office := queryParams.Get("office")
	order := queryParams.Get("order")

	attendances, err := ah.attService.GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, office, order)
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
	order := queryParams.Get("order")
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

	attCreate, err := ah.attService.CreateAttendance(user.ID, input)
	switch {
	case err != nil: // error internal server
		response, _ := json.Marshal(utils.APIResponse(err.Error(), http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		formatter := attendanceResponse.FormatAtt(attCreate)
		response, _ := json.Marshal(utils.APIResponse("Success Create Attendace Data", http.StatusOK, true, formatter))

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

	attUpdate, err := ah.attService.UpdateAttendance(user.ID, input)
	switch {
	case err != nil:
		response, _ := json.Marshal(utils.APIResponse(err.Error(), http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		formatter := attendanceResponse.FormatUpdateAtt(attUpdate)
		response, _ := json.Marshal(utils.APIResponse("Success Update Day Data", http.StatusOK, true, formatter))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
