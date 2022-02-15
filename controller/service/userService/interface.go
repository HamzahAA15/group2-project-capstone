package userService

import "sirclo/project-capstone/entities/userEntities"

type UserServiceInterface interface {
	GetUsers() ([]userEntities.User, error)
	GetUser(id string) (userEntities.User, error)
	DeleteUser(loginId string) error
}
