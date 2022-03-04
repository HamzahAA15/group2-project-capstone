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
	return users, err
}

func (us *userService) GetUser(id string) (userEntities.User, error) {
	user, err := us.userRepository.GetUser(id)
	return user, err
}

func (us *userService) CreateUser(input userRequest.CreateUserInput) (userEntities.User, error) {
	user := userEntities.User{}
	user.ID = uuid.New().String()
	user.Avatar = fmt.Sprintf("https://ui-avatars.com/api/?name=%s", input.Name)
	user.Nik = input.Nik
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

	user.Nik = input.Nik
	user.Email = input.Email
	passwordHash, _ := validation.HashPassword(input.Password)
	user.Password = passwordHash
	user.Name = input.Name
	user.Phone = input.Phone
	user.UpdatedAt = time.Now()

	updateUser, err := us.userRepository.UpdateUser(user)
	return updateUser, err
}

func (us *userService) UploadAvatarUser(id string, imageURL string) error {
	user, errGetUser := us.GetUser(id)
	if errGetUser != nil {
		return errGetUser
	}
	user.ID = id
	user.Avatar = imageURL
	user.UpdatedAt = time.Now()

	err := us.userRepository.UploadAvatarUser(user)
	return err
}
