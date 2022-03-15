package userService

import (
	"errors"
	"fmt"
	"testing"

	_userEntities "sirclo/project-capstone/entities/userEntities"
	_mockRepo "sirclo/project-capstone/mocks/user/repository"
	_utils "sirclo/project-capstone/utils"
	_userRequest "sirclo/project-capstone/utils/request/userRequest"
	_validation "sirclo/project-capstone/utils/validation"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	require "github.com/stretchr/testify/require"
)

func TestLoginUser(t *testing.T) {
	userData, password := randomUser(t)

	testCases := []struct {
		name     string
		body     _userRequest.LoginUserInput
		mockRepo func(user *_mockRepo.MockUserRepoInterface)
	}{
		{
			name: "success",
			body: _userRequest.LoginUserInput{
				Identity: userData.Email,
				Password: password,
			},
			mockRepo: func(user *_mockRepo.MockUserRepoInterface) {
				user.EXPECT().Login(gomock.Any()).Return(userData, nil).Times(1)
			},
		},
		{
			name: "error email",
			body: _userRequest.LoginUserInput{
				Identity: "error email",
				Password: password,
			},
			mockRepo: func(user *_mockRepo.MockUserRepoInterface) {
				user.EXPECT().Login(gomock.Any()).Return(userData, errors.New("error email")).Times(1)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockRepo)

			s := NewUserService(mockRepo)
			_, _ = s.LoginUserService(tc.body)
		})
	}
}

func TestGetUser(t *testing.T) {
	userData, _ := randomUser(t)

	testCases := []struct {
		name     string
		idUser   string
		mockRepo func(user *_mockRepo.MockUserRepoInterface)
	}{
		{
			name:   "success",
			idUser: "1",
			mockRepo: func(user *_mockRepo.MockUserRepoInterface) {
				user.EXPECT().GetUser(gomock.Eq("1")).Return(userData, nil).Times(1)
			},
		},
		{
			name:   "error",
			idUser: "2",
			mockRepo: func(user *_mockRepo.MockUserRepoInterface) {
				user.EXPECT().GetUser(gomock.Eq("2")).Return(userData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockRepo)

			s := NewUserService(mockRepo)
			_, _ = s.GetUser(tc.idUser)
		})
	}
}

func TestCreateUser(t *testing.T) {
	userData, password := randomUser(t)

	testCases := []struct {
		name     string
		body     _userRequest.CreateUserInput
		mockRepo func(user *_mockRepo.MockUserRepoInterface)
	}{
		{
			name: "success",
			body: _userRequest.CreateUserInput{
				Email:    userData.Email,
				Password: password,
			},
			mockRepo: func(user *_mockRepo.MockUserRepoInterface) {
				user.EXPECT().CreateUser(gomock.Any()).Times(1).Return(userData, nil)
			},
		},
		{
			name: "error",
			body: _userRequest.CreateUserInput{
				Email:    "email",
				Password: "password",
			},
			mockRepo: func(user *_mockRepo.MockUserRepoInterface) {
				user.EXPECT().CreateUser(gomock.Any()).Times(1).Return(userData, errors.New("error"))
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockRepo)

			s := NewUserService(mockRepo)
			_, _ = s.CreateUser(tc.body)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	userData, password := randomUser(t)

	testCases := []struct {
		name     string
		idUser   string
		body     _userRequest.UpdateUserInput
		mockRepo func(user *_mockRepo.MockUserRepoInterface)
	}{
		{
			name:   "success",
			idUser: "1",
			body: _userRequest.UpdateUserInput{
				Password: password,
			},
			mockRepo: func(user *_mockRepo.MockUserRepoInterface) {
				user.EXPECT().GetUser(gomock.Eq("1")).Return(userData, nil).Times(1)
				user.EXPECT().UpdateUser(gomock.Any()).Return(userData, nil).Times(1)
			},
		},
		{
			name:   "Error UserID",
			idUser: "2",
			body: _userRequest.UpdateUserInput{
				Password: password,
			},
			mockRepo: func(user *_mockRepo.MockUserRepoInterface) {
				user.EXPECT().GetUser(gomock.Eq("2")).Return(userData, errors.New("error")).Times(1)
			},
		},
		{
			name:   "Error Update Data",
			idUser: "3",
			body: _userRequest.UpdateUserInput{
				Password: password,
			},
			mockRepo: func(user *_mockRepo.MockUserRepoInterface) {
				user.EXPECT().GetUser(gomock.Eq("3")).Return(userData, nil).Times(1)
				user.EXPECT().UpdateUser(gomock.Any()).Return(userData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockRepo)

			s := NewUserService(mockRepo)
			_, _ = s.UpdateUser(tc.idUser, tc.body)
		})
	}
}

func TestUploadAvatar(t *testing.T) {
	userData, _ := randomUser(t)

	testCases := []struct {
		name     string
		idUser   string
		body     _userEntities.User
		mockRepo func(user *_mockRepo.MockUserRepoInterface)
	}{
		{
			name:   "success",
			idUser: "1",
			body: _userEntities.User{
				Avatar: userData.Avatar,
			},
			mockRepo: func(user *_mockRepo.MockUserRepoInterface) {
				user.EXPECT().GetUser(gomock.Eq("1")).Return(userData, nil).Times(1)
				user.EXPECT().UploadAvatarUser(gomock.Any()).Return(nil).Times(1)
			},
		},
		{
			name:   "Error UserID",
			idUser: "2",
			body: _userEntities.User{
				Avatar: userData.Avatar,
			},
			mockRepo: func(user *_mockRepo.MockUserRepoInterface) {
				user.EXPECT().GetUser(gomock.Eq("2")).Return(userData, errors.New("error")).Times(1)
			},
		},
		{
			name:   "Error Upload Avatar",
			idUser: "3",
			body: _userEntities.User{
				Avatar: userData.Avatar,
			},
			mockRepo: func(user *_mockRepo.MockUserRepoInterface) {
				user.EXPECT().GetUser(gomock.Eq("3")).Return(userData, nil).Times(1)
				user.EXPECT().UploadAvatarUser(gomock.Any()).Return(errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockRepo)

			s := NewUserService(mockRepo)
			_ = s.UploadAvatarUser(tc.idUser, tc.body.Avatar)
		})
	}
}

func randomUser(t *testing.T) (user _userEntities.User, password string) {
	password = _utils.RandomString(6)
	hashedPassword, err := _validation.HashPassword(password)
	require.NoError(t, err)

	user = _userEntities.User{
		ID:       uuid.New().String(),
		Avatar:   fmt.Sprintf("https://ui-avatars.com/api/?name=%s", _utils.RandomString(5)),
		Name:     _utils.RandomString(5),
		Email:    fmt.Sprintf(_utils.RandomString(6) + "@" + _utils.RandomString(4) + ".com"),
		Password: hashedPassword,
	}

	return
}
