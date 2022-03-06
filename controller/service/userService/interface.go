package userService

import (
	"sirclo/project-capstone/entities/userEntities"
	"sirclo/project-capstone/utils/request/userRequest"
)

type UserServiceInterface interface {
	LoginUserService(input userRequest.LoginUserInput) (userEntities.User, error)
	GetUser(id string) (userEntities.User, error)
	CreateUser(input userRequest.CreateUserInput) (userEntities.User, error)
	UpdateUser(id string, input userRequest.UpdateUserInput) (userEntities.User, error)
	UploadAvatarUser(id string, imageURL string) error
}
