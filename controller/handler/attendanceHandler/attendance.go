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
	employee := queryParams["employee"]
	time := queryParams["time"]

	attendances, err := ah.attService.GetAttendances(employee[0], time[0])
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

func (ah *attHandler) CreateAttendance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForContext(ctx)

	var input attendanceRequest.CreateAttRequest
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&input)

	attCreate, err := ah.attService.CreateAttendance(user.ID, input)
	switch {
	case err != nil: // error internal server
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

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

	checkUserRole := ah.attService.CheckUserRole(user.ID)
	if checkUserRole != "admin" {
		response, _ := json.Marshal(utils.APIResponse("You are not admin", http.StatusUnauthorized, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
		return
	}

	var input attendanceRequest.UpdateAttRequest
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&input)

	attUpdate, err := ah.attService.UpdateAttendance(user.ID, input)
	switch {
	case err != nil:
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

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
