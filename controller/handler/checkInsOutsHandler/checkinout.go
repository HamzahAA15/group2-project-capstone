package checkInsOutsHandler

import (
	"encoding/json"
	"net/http"
	"sirclo/project-capstone/controller/service/checkInsOutsService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/request/checkInsOutsRequest"
	"sirclo/project-capstone/utils/response/checkinsoutsResponse"
)

type checkInOutHandler struct {
	checkInOutService checkInsOutsService.CheckinoutServiceInterface
}

func NewCheckInOutHandler(checkInOutService checkInsOutsService.CheckinoutServiceInterface) CheckInOutHandlerInterface {
	return &checkInOutHandler{
		checkInOutService: checkInOutService,
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
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

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

func (ch *checkInOutHandler) GetsByUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := middleware.ForContext(ctx).ID

	CheckInsOuts, err := ch.checkInOutService.GetByUser(userID)
	switch {
	case err != nil:
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		var data []checkinsoutsResponse.CheckInsOutsResponse
		for _, val := range CheckInsOuts {
			dayFormatter := checkinsoutsResponse.FormatCheckInsOuts(val)
			data = append(data, dayFormatter)
		}

		response, _ := json.Marshal(utils.APIResponse("Success Get Data", http.StatusOK, true, data))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (ch *checkInOutHandler) CheckinsHandler(w http.ResponseWriter, r *http.Request) {
	var input checkInsOutsRequest.CheckInsRequest
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&input)

	checkIns, err := ch.checkInOutService.Checkin(input)
	switch {
	case err != nil: // error internal server
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		formatter := checkinsoutsResponse.FormatCheckInsOuts(checkIns)
		response, _ := json.Marshal(utils.APIResponse("Check-Ins Success", http.StatusOK, true, formatter))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
