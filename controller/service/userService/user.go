package userService

import (
	"fmt"
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

func (us *userService) LoginUserService(input userRequest.LoginUserInput) (userEntities.User, error) {
	email := input.Identity
	password := input.Password

	var user userEntities.User
	user, err := us.userRepository.Login(email)
	if err != nil {
		return user, err
	}

	err = validation.CheckPassword(password, user.Password)
	return user, err
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
	user.Avatar = "https://drive.google.com/file/d/1LUJXozBG_pAiGHNVM2tYgYyYxUH3knf-"
	user.Username = input.Username
	user.Email = input.Email
	passwordHash, _ := validation.HashPassword(input.Password)
	user.Password = passwordHash
	user.Name = input.Name
	user.Phone = input.Phone
	user.Role = input.Role
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
	user.Name = input.Name
	user.Phone = input.Phone
	user.UpdatedAt = time.Now()

	updateUser, err := us.userRepository.UpdateUser(user)
	return updateUser, err
}

func (us *userService) DeleteUser(loginId string) error {
	err := us.userRepository.DeleteUser(loginId)
	return err
}

func (us *userService) UploadAvatarUser(id string, imageURL string) error {
	user, err1 := us.GetUser(id)
	fmt.Println("service upload: ", user)
	fmt.Println("service err_upload: ", err1)
	if err1 != nil {
		return err1
	}
	user.ID = id
	user.Avatar = imageURL
	user.UpdatedAt = time.Now()

	err := us.userRepository.UploadAvatarUser(user)
	return err
}
