package userService

import (
	"sirclo/project-capstone/entities/userEntities"
	"sirclo/project-capstone/repository/userRepository"
)

type userService struct {
	userRepository userRepository.UserRepoInterface
}

func NewUserService(repo userRepository.UserRepoInterface) UserServiceInterface {
	return &userService{
		userRepository: repo,
	}
}

func (us *userService) GetUsers() ([]userEntities.User, error) {
	users, err := us.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *userService) GetUser(id string) (userEntities.User, error) {
	user, err := us.userRepository.GetUser(id)
	return user, err
}

func (us *userService) DeleteUser(loginId string) error {
	err := us.userRepository.DeleteUser(loginId)
	return err
}
