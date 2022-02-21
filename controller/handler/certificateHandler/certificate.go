package certificateHandler

import (
	"encoding/json"
	"net/http"
	"sirclo/project-capstone/controller/service/certificateService"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/response/certificateResponse"
)

type certificateHandler struct {
	certificateService certificateService.CertificateServiceInterface
}

func NewCertificateHandler(dayService certificateService.CertificateServiceInterface) CertificateHandlerInterface {
	return &certificateHandler{
		certificateService: dayService,
	}
}

func (ch *certificateHandler) GetCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	certificates, err := ch.certificateService.GetCertificates()
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
