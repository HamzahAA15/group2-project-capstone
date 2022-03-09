package userService

import (
	"errors"
	"sirclo/project-capstone/entities/userEntities"
	"sirclo/project-capstone/utils/request/userRequest"
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

// func TestLoginUser(t *testing.T) {
// 	userServiceMock := NewUserService(mockUserRepository{})
// 	input := userRequest.LoginUserInput{
// 		Identity: "felicia@sirclo.com",
// 		Password: "123456789qwerty",
// 	}
// 	user, err := userServiceMock.LoginUserService(input)
// 	// expected := userEntities.User{
// 	// 	ID:       "1",
// 	// 	Email:    "felicia@sirclo.com",
// 	// 	Password: "123456789",
// 	// 	Role:     "admin",
// 	// }
// 	assert.Nil(t, err)
// 	assert.Equal(t, "1", user.ID)
// }

// func TestUpdateUser(id string, input userRequest.UpdateUserInput) (userEntities.User, error) {

// }

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
	return userEntities.User{
		ID:       "1",
		Email:    "felicia@sirclo.com",
		Password: "123456789qwerty",
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
	return userEntities.User{}, nil

}
func (m mockUserRepository) UploadAvatarUser(user userEntities.User) error {
	return nil
}
