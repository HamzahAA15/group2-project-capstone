package userService

import (
	"sirclo/project-capstone/entities/userEntities"
	"sirclo/project-capstone/utils/request/userRequest"
)

//go:generate mockgen --destination=./../../../mocks/user/service/mock_service_user.go -source=interface.go
type UserServiceInterface interface {
	LoginUserService(input userRequest.LoginUserInput) (userEntities.User, error)
	GetUser(id string) (userEntities.User, error)
	CreateUser(input userRequest.CreateUserInput) (userEntities.User, error)
	UpdateUser(id string, input userRequest.UpdateUserInput) (userEntities.User, error)
	UploadAvatarUser(id string, imageURL string) error
}
