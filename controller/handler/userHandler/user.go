package userHandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/middleware"
	"sirclo/project-capstone/utils/request/userRequest"
	"sirclo/project-capstone/utils/upload"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gorilla/mux"
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
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error"))
	}

	user, err_login := uh.userService.LoginUserService(input)
	if err_login != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
	}
	token, err_token := middleware.GenerateToken(user.ID)
	if err_token != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("failed to generate token"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))

}

func (uh *userHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userService.GetUsers()
	switch {
	case err != nil: // error internal server
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		response, _ := json.Marshal(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(users)
		w.Write(response)
	}
}

func (uh *userHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := uh.userService.GetUser(id)
	switch {
	case err == sql.ErrNoRows: //check data is null?
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		response, _ := json.Marshal(http.StatusNotFound)
		w.Write(response)
	case err != nil: // error internal server
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		response, _ := json.Marshal(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(user)
		w.Write(response)
	}
}

func (uh *userHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input userRequest.CreateUserInput
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&input)

	user, err := uh.userService.CreateUser(input)
	switch {
	case err != nil: // error internal server
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		response, _ := json.Marshal(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(user)
		w.Write(response)
	}
}

func (uh *userHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForContext(ctx)

	var input userRequest.UpdateUserInput

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&input)

	userUpdate, err := uh.userService.UpdateUser(user.ID, input)
	switch {
	case err != nil: // error internal server
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		response, _ := json.Marshal(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(userUpdate)
		w.Write(response)
	}
}

func (uh *userHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForContext(ctx)

	err := uh.userService.DeleteUser(user.ID)
	switch {
	case err != nil: // error internal server
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		response, _ := json.Marshal(http.StatusInternalServerError)
		w.Write(response)
	default: // default response success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(http.StatusOK)
		w.Write(response)
	}
}

func (uh *userHandler) UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForContext(ctx)
	maxSize := int64(2048000)

	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(fmt.Sprintf("Image too large. Max Size: %v", maxSize)))
		return
	}

	file, fileHeader, err := r.FormFile("avatar")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not get uploaded file"))
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		response, _ := json.Marshal(http.StatusBadGateway)
		w.Write(response)
		return
	}

	fileLoc, err := upload.UploadFile(user.ID, "users", s, file, fileHeader)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("file extension isn't equal to .png, .jpg. and .jpeg"))
		return
	}

	err_upload := uh.userService.UploadAvatarUser(user.ID, fileLoc)
	if err_upload != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to upload photo"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Image uploaded successfully: %v", fileLoc)))
}
