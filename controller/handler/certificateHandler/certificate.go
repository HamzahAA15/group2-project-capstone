package certificateHandler

import (
	"encoding/json"
	"net/http"
	"os"
	"sirclo/project-capstone/controller/service/certificateService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/request/certificateRequest"
	"sirclo/project-capstone/utils/response/certificateResponse"
	"sirclo/project-capstone/utils/upload"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
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

func (ch *certificateHandler) UploadCertificateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForContext(ctx)
	maxSize := int64(2048000)

	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		response, _ := json.Marshal(utils.APIResponse("Image too large. Max Size", http.StatusUnprocessableEntity, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(response)
		return
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		response, _ := json.Marshal(utils.APIResponse("Could not get uploaded file", http.StatusBadRequest, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	defer file.Close()

	s, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("KEYID"),
			os.Getenv("SECRETKEY"),
			""),
	})
	if err != nil {
		response, _ := json.Marshal(utils.APIResponse("Could not uploaded file", http.StatusBadRequest, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	fileLoc, err := upload.UploadFile(user.ID+time.Now().Format("20060102150405"), "certificates", s, file, fileHeader)
	if err != nil {
		response, _ := json.Marshal(utils.APIResponse("file extension isn't equal to .png, .jpg. and .jpeg", http.StatusBadRequest, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	countDosage := ch.certificateService.GetVaccineDose(user.ID) + 1
	countVaccine := ch.certificateService.CountVaccineAccept(user.ID, countDosage)
	if countVaccine > 0 {
		response, _ := json.Marshal(utils.APIResponse("Please wait till your vaccine certificate has been verified by Admin.", http.StatusBadRequest, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	input := certificateRequest.CertificateUploadRequest{
		Image:  fileLoc,
		Dosage: countDosage,
	}

	err_upload := ch.certificateService.UploadCertificateVaccine(user.ID, input)
	if err_upload != nil {
		response, _ := json.Marshal(utils.APIResponse("failed to upload image", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	response, _ := json.Marshal(utils.APIResponse("Image uploaded successfully", http.StatusOK, true, nil))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
