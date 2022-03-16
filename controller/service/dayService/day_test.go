package dayService

import (
	"errors"
	"testing"

	_dayEntities "sirclo/project-capstone/entities/dayEntities"
	_mockDayRepo "sirclo/project-capstone/mocks/day/repository"
	_mockUserRepo "sirclo/project-capstone/mocks/user/repository"
	_dayRequest "sirclo/project-capstone/utils/request/dayRequest"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

func TestGetDays(t *testing.T) {
	daysData := dummyDays(t)

	testCases := []struct {
		name     string
		idOffice string
		date     string
		mockRepo func(day *_mockDayRepo.MockDayRepoInterface, user *_mockUserRepo.MockUserRepoInterface)
	}{
		{
			name:     "success",
			idOffice: "1",
			date:     "2001-01-01",
			mockRepo: func(day *_mockDayRepo.MockDayRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				day.EXPECT().GetDays(gomock.Eq("1"), gomock.Eq("2001-01-01")).Return(daysData, nil).Times(1)
			},
		},
		{
			name:     "error",
			idOffice: "2",
			date:     "2002-02-02",
			mockRepo: func(day *_mockDayRepo.MockDayRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				day.EXPECT().GetDays(gomock.Eq("2"), gomock.Eq("2002-02-02")).Return(daysData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDayRepo := _mockDayRepo.NewMockDayRepoInterface(ctrl)
			mockUserRepo := _mockUserRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockDayRepo, mockUserRepo)

			s := NewDayService(mockDayRepo, mockUserRepo)
			_, _ = s.GetDays(tc.idOffice, tc.date)
		})
	}
}

func TestGetDaysID(t *testing.T) {
	dayData := dummyDay(t)

	testCases := []struct {
		name     string
		idDay    string
		mockRepo func(day *_mockDayRepo.MockDayRepoInterface, user *_mockUserRepo.MockUserRepoInterface)
	}{
		{
			name:  "success",
			idDay: "1",
			mockRepo: func(day *_mockDayRepo.MockDayRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				day.EXPECT().GetDayID(gomock.Eq("1")).Return(dayData, nil).Times(1)
			},
		},
		{
			name:  "error",
			idDay: "2",
			mockRepo: func(day *_mockDayRepo.MockDayRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				day.EXPECT().GetDayID(gomock.Eq("2")).Return(dayData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDayRepo := _mockDayRepo.NewMockDayRepoInterface(ctrl)
			mockUserRepo := _mockUserRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockDayRepo, mockUserRepo)

			s := NewDayService(mockDayRepo, mockUserRepo)
			_, _ = s.GetDaysID(tc.idDay)
		})
	}
}

func TestUpdateDays(t *testing.T) {
	testCases := []struct {
		name     string
		body     _dayRequest.DayUpdateRequest
		mockRepo func(day *_mockDayRepo.MockDayRepoInterface, user *_mockUserRepo.MockUserRepoInterface)
	}{
		{
			name: "success",
			body: _dayRequest.DayUpdateRequest{
				ID:    uuid.New().String(),
				Quota: 50,
			},
			mockRepo: func(day *_mockDayRepo.MockDayRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				day.EXPECT().UpdateDay(gomock.Any()).Return(_dayEntities.Day{}, nil).Times(1)
			},
		},
		{
			name: "error",
			body: _dayRequest.DayUpdateRequest{
				ID:    uuid.New().String(),
				Quota: 100,
			},
			mockRepo: func(day *_mockDayRepo.MockDayRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				day.EXPECT().UpdateDay(gomock.Any()).Return(_dayEntities.Day{}, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDayRepo := _mockDayRepo.NewMockDayRepoInterface(ctrl)
			mockUserRepo := _mockUserRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockDayRepo, mockUserRepo)

			s := NewDayService(mockDayRepo, mockUserRepo)
			_, _ = s.UpdateDays(tc.body)
		})
	}
}

func dummyDays(t *testing.T) (days []_dayEntities.Day) {
	days = []_dayEntities.Day{
		{
			ID:    uuid.New().String(),
			Quota: 50,
		},
	}

	return
}

func dummyDay(t *testing.T) (day _dayEntities.Day) {
	day = _dayEntities.Day{
		ID:    uuid.New().String(),
		Quota: 50,
	}

	return
}
