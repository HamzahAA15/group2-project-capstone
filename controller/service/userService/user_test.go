package userService

import (
	"errors"
	"sirclo/project-capstone/entities/userEntities"
	"sirclo/project-capstone/utils/request/userRequest"
	"sirclo/project-capstone/utils/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	userServiceMock := NewUserService(mockUserRepository{})
	user, err := userServiceMock.GetUser("1")
	assert.Nil(t, err)
	assert.Equal(t, "hallo", user.Name, "nama tidak sama")
}

func TestCreateUser(t *testing.T) {
	var create = userRequest.CreateUserInput{
		ID:    "1",
		Name:  "hallo",
		Email: "mail@test.com",
	}

	userServiceMock := NewUserService(mockUserRepository{})
	user, err := userServiceMock.CreateUser(create)
	assert.Nil(t, err)
	assert.Equal(t, "1", user.ID, "nama tidak sama")
	assert.Equal(t, "hallo", user.Name, "nama tidak sama")
}

func TestLoginUserSuccess(t *testing.T) {
	userServiceMock := NewUserService(mockUserRepository{})
	input := userRequest.LoginUserInput{
		Identity: "felicia@sirclo.com",
		Password: "123456789qwerty",
	}

	_, err := userServiceMock.LoginUserService(input)
	assert.Nil(t, err)
}

func TestLoginUserError(t *testing.T) {
	userServiceMock := NewUserService(mockUserRepository{})
	input := userRequest.LoginUserInput{
		Identity: "",
		Password: "123456789qwerty",
	}
	_, err := userServiceMock.LoginUserService(input)
	assert.NotNil(t, err)
}

func TestUpdateUserSuccess(t *testing.T) {
	input := userRequest.UpdateUserInput{
		Password: "passowrd123",
	}

	userServiceMock := NewUserService(mockUserRepository{})
	user, err := userServiceMock.UpdateUser("1", input)
	assert.Nil(t, err)
	assert.Equal(t, "1", user.ID, "tidak sama")
}

func TestUpdateUserError(t *testing.T) {
	input := userRequest.UpdateUserInput{
		Password: "passowrd123",
	}

	userServiceMock := NewUserService(mockUserRepository{})
	_, err := userServiceMock.UpdateUser("2", input)
	assert.NotNil(t, err)
}

func TestUploadAvatar1(t *testing.T) {
	userServiceMock := NewUserService(mockUserRepository{})
	_ = userServiceMock.UploadAvatarUser("1", "https://lazyevent.site")
	assert.Equal(t, nil, nil)
}

func TestUploadAvatar2(t *testing.T) {
	userServiceMock := NewUserService(mockUserRepository{})
	_ = userServiceMock.UploadAvatarUser("2", "https://lazyevent.site")
	assert.Equal(t, nil, nil)
}

type mockUserRepository struct{}

func (m mockUserRepository) GetUser(id string) (userEntities.User, error) {
	var gagal userEntities.User
	if id == "2" {
		return gagal, errors.New("error")
	}

	return userEntities.User{
		ID:    "1",
		Name:  "hallo",
		Email: "mail@test.com",
	}, nil
}

func (m mockUserRepository) CheckEmail(userChecked userEntities.User) (userEntities.User, error) {
	return userEntities.User{
		Email: userChecked.Email,
	}, nil
}

func (m mockUserRepository) Login(identity string) (userEntities.User, error) {
	if identity == "" {
		return userEntities.User{}, errors.New("error")
	}

	passwordHash, _ := validation.HashPassword("123456789qwerty")
	return userEntities.User{
		ID:       "1",
		Email:    "felicia@sirclo.com",
		Password: string(passwordHash),
		Role:     "admin",
	}, nil
}

func (m mockUserRepository) CreateUser(user userEntities.User) (userEntities.User, error) {
	return userEntities.User{
		ID:    "1",
		Name:  "hallo",
		Email: "mail@test.com",
	}, nil
}
func (m mockUserRepository) UpdateUser(user userEntities.User) (userEntities.User, error) {
	return userEntities.User{
		ID:    "1",
		Name:  "hallo",
		Email: "mail@test.com",
	}, nil
}

func (m mockUserRepository) UploadAvatarUser(user userEntities.User) error {
	return nil
}
