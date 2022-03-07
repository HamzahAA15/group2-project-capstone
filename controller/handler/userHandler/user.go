package userHandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/utils"
	"sirclo/project-capstone/utils/request/userRequest"
	"sirclo/project-capstone/utils/response/userResponse"
	"sirclo/project-capstone/utils/upload"
	"sirclo/project-capstone/utils/validation"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type userHandler struct {
	userService userService.UserServiceInterface
}

func NewUserHandler(userService userService.UserServiceInterface) UserHandlerInterface {
	return &userHandler{
		userService: userService,
	}
}

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input userRequest.LoginUserInput
	json.NewDecoder(r.Body).Decode(&input)

	errValidation := validation.CheckEmpty(input.Identity, input.Password)
	if errValidation != nil {
		response, _ := json.Marshal(utils.APIResponse(errValidation.Error(), http.StatusUnprocessableEntity, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(response)
		return
	}

	user, err_login := uh.userService.LoginUserService(input)
	if err_login != nil {
		response, _ := json.Marshal(utils.APIResponse("Login Failed", http.StatusUnprocessableEntity, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(response)
		return
	}
	token, err_token := middleware.GenerateToken(user.ID)
	if err_token != nil {
		response, _ := json.Marshal(utils.APIResponse("Failed to Generate Token", http.StatusBadRequest, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	formatter := userResponse.FormatAuth(token, user.Role)
	response, _ := json.Marshal(utils.APIResponse("Success Generate Token", http.StatusOK, true, formatter))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (uh *userHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForContext(ctx)

	getUser, err := uh.userService.GetUser(user.ID)
	switch {
	case err == sql.ErrNoRows: //check data is null?
		response, _ := json.Marshal(utils.APIResponse("Data Not Found", http.StatusNotFound, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(response)
	case err != nil: // error internal server
		response, _ := json.Marshal(utils.APIResponse("Something Went Wrong", http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		formatter := userResponse.FormatUser(getUser)
		response, _ := json.Marshal(utils.APIResponse("Success Get User By ID", http.StatusOK, true, formatter))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (uh *userHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForContext(ctx)

	getUser, _ := uh.userService.GetUser(user.ID)
	if getUser.Role != "admin" {
		response, _ := json.Marshal(utils.APIResponse("You don't have permission to access it", http.StatusUnauthorized, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
		return
	}

	var input userRequest.CreateUserInput
	json.NewDecoder(r.Body).Decode(&input)

	errValidation := validation.CheckEmpty(input.OfficeID, input.Nik, input.Email, input.Password, input.Name, input.Phone, input.Role)
	if errValidation != nil {
		response, _ := json.Marshal(utils.APIResponse(errValidation.Error(), http.StatusUnprocessableEntity, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(response)
		return
	}

	_, err := uh.userService.CreateUser(input)
	switch {
	case err != nil: // error internal server
		response, _ := json.Marshal(utils.APIResponse(err.Error(), http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		response, _ := json.Marshal(utils.APIResponse("Success Create User Data", http.StatusOK, true, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (uh *userHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForContext(ctx)

	var input userRequest.UpdateUserInput
	json.NewDecoder(r.Body).Decode(&input)

	_, err := uh.userService.UpdateUser(user.ID, input)
	switch {
	case err != nil: // error internal server
		response, _ := json.Marshal(utils.APIResponse(err.Error(), http.StatusInternalServerError, false, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		response, _ := json.Marshal(utils.APIResponse("Success Update User Data", http.StatusOK, true, nil))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (uh *userHandler) UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForContext(ctx)
	maxSize := int64(5120000)

	file, fileHeader, err := r.FormFile("avatar")
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

		fileLoc, errExtension := upload.UploadFile(user.ID, "users", s, file, fileHeader)
		if errExtension != nil {
			response, _ := json.Marshal(utils.APIResponse("file extension isn't equal to .png, .jpg. and .jpeg", http.StatusBadRequest, false, nil))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

		err_upload := uh.userService.UploadAvatarUser(user.ID, fileLoc)
		if err_upload != nil {
			response, _ := json.Marshal(utils.APIResponse("failed to upload photo", http.StatusInternalServerError, false, nil))

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
