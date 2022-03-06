package logcatHandler

import (
	"encoding/json"
	"net/http"
	"sirclo/project-capstone/controller/service/logcatService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/response/logcatResponse"
)

type logcatHandler struct {
	userService   userService.UserServiceInterface
	logcatService logcatService.LogcatServiceInterface
}

func NewLogcatHandler(userService userService.UserServiceInterface, logcatService logcatService.LogcatServiceInterface) LogcatHandlerInterface {
	return &logcatHandler{
		userService:   userService,
		logcatService: logcatService,
	}
}

func (lh *logcatHandler) LogcatsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := middleware.ForContext(ctx).ID

	getUser, _ := lh.userService.GetUser(userID)
	if getUser.Role != "admin" {
		response, _ := json.Marshal(utils.APIResponse("You are not admin", http.StatusUnauthorized, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
		return
	}

	logcats, err := lh.logcatService.GetLogcats()
	switch {
	case err != nil: // error internal server
		response, _ := json.Marshal(utils.APIResponse("Something Went Wrong", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		var data []logcatResponse.LogcatResponse
		for i := 0; i < len(logcats); i++ {
			formatter := logcatResponse.FormatLog(logcats[i])
			data = append(data, formatter)
		}

		response, _ := json.Marshal(utils.APIResponse("Success Get Log Data", http.StatusOK, true, data))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (lh *logcatHandler) LogcatUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := middleware.ForContext(ctx).ID

	logcats, err := lh.logcatService.GetLogcatUser(userID)
	switch {
	case err != nil: // error internal server
		response, _ := json.Marshal(utils.APIResponse("Something Went Wrong", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		var data []logcatResponse.LogcatResponse
		for i := 0; i < len(logcats); i++ {
			formatter := logcatResponse.FormatLog(logcats[i])
			data = append(data, formatter)
		}

		response, _ := json.Marshal(utils.APIResponse("Success Get Own Log Data", http.StatusOK, true, data))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
