package userRepository

import (
	"sirclo/project-capstone/entities/userEntities"
)

//go:generate mockgen --destination=./../../mocks/user/repository/mock_repository_user.go -source=interface.go
type UserRepoInterface interface {
	CheckEmail(userChecked userEntities.User) (userEntities.User, error)
	Login(identity string) (userEntities.User, error)
	GetUser(id string) (userEntities.User, error)
	CreateUser(user userEntities.User) (userEntities.User, error)
	UpdateUser(user userEntities.User) (userEntities.User, error)
	UploadAvatarUser(user userEntities.User) error
}
