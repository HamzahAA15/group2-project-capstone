package userHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sirclo/project-capstone/controller/service/userService"
	"sirclo/project-capstone/utils/request/userRequest"

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

func (uh *userHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["id"]

	users, err := uh.userService.GetUsers()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("not found"))
	}

	// var data []userEntities.User
	// for i, _ := range users {
	// 	data = append(data, users[i])
	// }
	response, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (uh *userHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input userRequest.CreateUserInput
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&input)

	user, err := uh.userService.CreateUser(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal service error"))
	}

	response, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (uh *userHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var input userRequest.UpdateUserInput

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&input)

	user, err := uh.userService.UpdateUser(id, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal service error"))
	}

	response, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (uh *userHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := uh.userService.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("not found"))
	}

	response, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (uh *userHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := uh.userService.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("not found"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("delete success"))
}
