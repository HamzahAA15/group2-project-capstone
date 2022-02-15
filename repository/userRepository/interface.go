package userRepository

import (
	"sirclo/project-capstone/entities/userEntities"
)

type UserRepoInterface interface {
	GetUsers() ([]userEntities.User, error)
	GetUser(id string) (userEntities.User, error)
	DeleteUser(loginId string) error
}
