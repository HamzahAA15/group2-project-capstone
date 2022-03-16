package logcatService

import (
	"errors"
	"testing"
	"time"

	_logcatEntities "sirclo/project-capstone/entities/logcatEntities"
	_mockRepo "sirclo/project-capstone/mocks/logcat/repository"
	_utils "sirclo/project-capstone/utils"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

func TestCreateLogcat(t *testing.T) {
	logcatData := randomLogcat(t)

	testCases := []struct {
		name           string
		body           _logcatEntities.Logcat
		mockLogcatRepo func(logcat *_mockRepo.MockLogcatRepoInterface)
	}{
		{
			name: "success",
			body: _logcatEntities.Logcat{
				ID:       logcatData.ID,
				Message:  logcatData.Message,
				Category: logcatData.Category,
			},
			mockLogcatRepo: func(logcat *_mockRepo.MockLogcatRepoInterface) {
				logcat.EXPECT().CreateLogcat(gomock.Any()).Return(logcatData, nil).Times(1)
			},
		},
		{
			name: "error",
			body: _logcatEntities.Logcat{
				ID:       logcatData.ID,
				Message:  logcatData.Message,
				Category: logcatData.Category,
			},
			mockLogcatRepo: func(logcat *_mockRepo.MockLogcatRepoInterface) {
				logcat.EXPECT().CreateLogcat(gomock.Any()).Return(logcatData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockRepo.NewMockLogcatRepoInterface(ctrl)
			tc.mockLogcatRepo(mockRepo)

			s := NewLogcatService(mockRepo)
			_, _ = s.CreateLogcat(tc.body.ID, tc.body.Message, tc.body.Category)
		})
	}
}

func TestGetLogcats(t *testing.T) {
	logcatData := randomLogcats(t)

	testCases := []struct {
		name           string
		mockLogcatRepo func(logcat *_mockRepo.MockLogcatRepoInterface)
	}{
		{
			name: "success",
			mockLogcatRepo: func(logcat *_mockRepo.MockLogcatRepoInterface) {
				logcat.EXPECT().GetLogcats().Return(logcatData, nil).Times(1)
			},
		},
		{
			name: "error",
			mockLogcatRepo: func(logcat *_mockRepo.MockLogcatRepoInterface) {
				logcat.EXPECT().GetLogcats().Return(logcatData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockRepo.NewMockLogcatRepoInterface(ctrl)
			tc.mockLogcatRepo(mockRepo)

			s := NewLogcatService(mockRepo)
			_, _ = s.GetLogcats()
		})
	}
}

func TestGetLogcatUser(t *testing.T) {
	logcatData := randomLogcats(t)

	testCases := []struct {
		name           string
		idUser         string
		mockLogcatRepo func(logcat *_mockRepo.MockLogcatRepoInterface)
	}{
		{
			name:   "success",
			idUser: "1",
			mockLogcatRepo: func(logcat *_mockRepo.MockLogcatRepoInterface) {
				logcat.EXPECT().GetLogcatUser(gomock.Eq("1")).Return(logcatData, nil).Times(1)
			},
		},
		{
			name:   "error",
			idUser: "2",
			mockLogcatRepo: func(logcat *_mockRepo.MockLogcatRepoInterface) {
				logcat.EXPECT().GetLogcatUser(gomock.Eq("2")).Return(logcatData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockRepo.NewMockLogcatRepoInterface(ctrl)
			tc.mockLogcatRepo(mockRepo)

			s := NewLogcatService(mockRepo)
			_, _ = s.GetLogcatUser(tc.idUser)
		})
	}
}

func randomLogcat(t *testing.T) (logcat _logcatEntities.Logcat) {
	logcat = _logcatEntities.Logcat{
		ID:        uuid.New().String(),
		Message:   _utils.RandomString(20),
		Category:  "logcat",
		UpdatedAt: time.Now(),
	}

	return
}

func randomLogcats(t *testing.T) (logcat []_logcatEntities.Logcat) {
	logcat = []_logcatEntities.Logcat{
		{
			ID:        uuid.New().String(),
			Message:   _utils.RandomString(20),
			Category:  "logcat",
			UpdatedAt: time.Now(),
		},
	}

	return
}
