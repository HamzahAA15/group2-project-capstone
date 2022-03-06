package certificateHandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sirclo/project-capstone/controller/service/certificateService"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/request/certificateRequest"
	"sirclo/project-capstone/utils/response/certificateResponse"
	"sirclo/project-capstone/utils/upload"
	"sirclo/project-capstone/utils/validation"
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

	queryParams := r.URL.Query()
	orderBy := queryParams.Get("order_by")

	user, _ := ch.userService.GetUser(userID)
	if user.Role != "admin" {
		response, _ := json.Marshal(utils.APIResponse("You don't have permission to access this data", http.StatusUnauthorized, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
		return
	}

	certificates, err := ch.certificateService.GetCertificates(orderBy)
	switch {
	case err != nil:
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		var data []certificateResponse.CertificatesResponse
		for _, val := range certificates {
			certificateFormatter := certificateResponse.FormatCertificates(val)
			data = append(data, certificateFormatter)
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
			certificateFormatter := certificateResponse.FormatCertificate(val)
			data = append(data, certificateFormatter)
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
	maxSize := int64(5120000)

	file, fileHeader, err := r.FormFile("image")
	switch err {
	case http.ErrMissingFile:
		response, _ := json.Marshal(utils.APIResponse("there are no files to upload", http.StatusUnprocessableEntity, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(response)
		return
	default:
		if fileHeader.Size > maxSize {
			response, _ := json.Marshal(utils.APIResponse(fmt.Sprintf("Image too large. Max Size: %v Kb", maxSize), http.StatusUnprocessableEntity, false, nil))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write(response)
			return
		}

		defer file.Close()

		s, _ := session.NewSession(&aws.Config{
			Region: aws.String(os.Getenv("REGION")),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("KEYID"),
				os.Getenv("SECRETKEY"),
				""),
		})

		fileLoc, errExtension := upload.UploadFile(user.ID+time.Now().Format("20060102150405"), "certificates", s, file, fileHeader)
		if errExtension != nil {
			response, _ := json.Marshal(utils.APIResponse("file extension isn't equal to .png, .jpg. and .jpeg", http.StatusBadRequest, false, nil))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

		countDosage := ch.certificateService.GetVaccineDose(user.ID, "approved") + 1
		countVaccine := ch.certificateService.CountVaccineIsPending(user.ID, countDosage)
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
}

func (ch *certificateHandler) VerifyCertificateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := middleware.ForContext(ctx).ID

	user, _ := ch.userService.GetUser(userID)
	if user.Role != "admin" {
		response, _ := json.Marshal(utils.APIResponse("You don't have permission to access this url", http.StatusUnauthorized, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
		return
	}

	var input certificateRequest.CertificateUploadRequest
	json.NewDecoder(r.Body).Decode(&input)

	dataVaccine, _ := ch.certificateService.GetCertificate(input.ID)
	if dataVaccine.Status != "pending" {
		response, _ := json.Marshal(utils.APIResponse("you cannot update this data again", http.StatusBadRequest, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	errValidation := validation.CheckEmpty(input.Status)
	if errValidation != nil {
		response, _ := json.Marshal(utils.APIResponse(errValidation.Error(), http.StatusUnprocessableEntity, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(response)
		return
	}

	_, err := ch.certificateService.VerifyCertificate(input.ID, userID, input)
	switch {
	case err == sql.ErrNoRows: //check data is null?
		response, _ := json.Marshal(utils.APIResponse("Data Not Found", http.StatusNotFound, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(response)
	case err != nil: // error internal server
		response, _ := json.Marshal(utils.APIResponse("Internal Server Error", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		response, _ := json.Marshal(utils.APIResponse("Success Update Data", http.StatusOK, true, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
