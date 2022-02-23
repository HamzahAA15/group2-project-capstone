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
		response, _ := json.Marshal(utils.APIResponse("Success Create User Data", http.StatusOK, true, formatter))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

}

func (ah *attHandler) UpdateAttendance(w http.ResponseWriter, r *http.Request) {

}
