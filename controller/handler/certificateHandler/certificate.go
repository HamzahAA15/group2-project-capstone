package certificateHandler

import (
	"encoding/json"
	"net/http"
	"sirclo/project-capstone/controller/service/certificateService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/response/certificateResponse"
)

type certificateHandler struct {
	certificateService certificateService.CertificateServiceInterface
	userService        userService.UserServiceInterface
}

func NewCertificateHandler(certificateService certificateService.CertificateServiceInterface, userService userService.UserServiceInterface) CertificateHandlerInterface {
	return &certificateHandler{
		certificateService: certificateService,
		userService:        userService,
	}
}

func (ch *certificateHandler) GetCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := middleware.ForContext(ctx).ID

	user, _ := ch.userService.GetUser(userID)
	if user.Role != "admin" {
		response, _ := json.Marshal(utils.APIResponse("You don't have permission to access this data", http.StatusUnauthorized, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
		return
	}

	certificates, err := ch.certificateService.GetCertificates(user.OfficeID)
	switch {
	case err != nil:
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		var data []certificateResponse.CertificateResponse
		for _, val := range certificates {
			dayFormatter := certificateResponse.FormatCertificate(val)
			data = append(data, dayFormatter)
		}

		response, _ := json.Marshal(utils.APIResponse("Success Get Certificates Data", http.StatusOK, true, data))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

}

func (ch *certificateHandler) GetCertificateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := middleware.ForContext(ctx).ID

	certificates, err := ch.certificateService.GetCertificateUser(userID)
	switch {
	case err != nil:
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		var data []certificateResponse.CertificateResponse
		for _, val := range certificates {
			dayFormatter := certificateResponse.FormatCertificate(val)
			data = append(data, dayFormatter)
		}

		response, _ := json.Marshal(utils.APIResponse("Success Get Certificates Data", http.StatusOK, true, data))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
