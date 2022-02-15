package userService

import (
	"sirclo/project-capstone/entities/userEntities"
	"sirclo/project-capstone/repository/userRepository"
	"sirclo/project-capstone/utils/request/userRequest"
	"sirclo/project-capstone/utils/validation"
	"time"

	"github.com/google/uuid"
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

func (us *userService) CreateUser(input userRequest.CreateUserInput) (userEntities.User, error) {
	user := userEntities.User{}
	user.ID = uuid.New().String()
	user.Username = input.Username
	user.Email = input.Email
	passwordHash, _ := validation.HashPassword(input.Password)
	user.Password = passwordHash
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	createUser, err := us.userRepository.CreateUser(user)
	return createUser, err
}

func (us *userService) UpdateUser(id string, input userRequest.UpdateUserInput) (userEntities.User, error) {
	user, err := us.GetUser(id)
	if err != nil {
		return user, err
	}

	user.Username = input.Username
	user.Email = input.Email
	passwordHash, _ := validation.HashPassword(input.Password)
	user.Password = passwordHash
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	updateUser, err := us.userRepository.UpdateUser(user)
	return updateUser, err
}

func (us *userService) DeleteUser(loginId string) error {
	err := us.userRepository.DeleteUser(loginId)
	return err
}
