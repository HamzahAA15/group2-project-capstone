package checkInsOutsService

import (
	"errors"
	"testing"
	"time"

	_attendanceEntities "sirclo/project-capstone/entities/attendanceEntities"
	_checkinsoutsEntities "sirclo/project-capstone/entities/checkinEntities"
	_mockCheckinsoutsRepo "sirclo/project-capstone/mocks/checkinout/repository"
	_utils "sirclo/project-capstone/utils"
	_checkInsOutsRequest "sirclo/project-capstone/utils/request/checkInsOutsRequest"

	gomock "github.com/golang/mock/gomock"
)

func TestGets(t *testing.T) {
	checkinsoutsData := dummyCheckinsouts(t)

	testCases := []struct {
		name                 string
		mockCheckinsoutsRepo func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface)
	}{
		{
			name: "success",
			mockCheckinsoutsRepo: func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface) {
				checkinsouts.EXPECT().Gets().Return(checkinsoutsData, nil).Times(1)
			},
		},
		{
			name: "error",
			mockCheckinsoutsRepo: func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface) {
				checkinsouts.EXPECT().Gets().Return(checkinsoutsData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockCheckinsoutsRepo.NewMockCheckInOutRepoInterface(ctrl)
			tc.mockCheckinsoutsRepo(mockRepo)

			s := NewCheckInOutService(mockRepo)
			_, _ = s.Gets()
		})
	}
}

func TestGetByUser(t *testing.T) {
	checkinsoutsData := dummyCheckinsouts(t)

	testCases := []struct {
		name                 string
		idUser               string
		mockCheckinsoutsRepo func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface)
	}{
		{
			name:   "success",
			idUser: "1",
			mockCheckinsoutsRepo: func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface) {
				checkinsouts.EXPECT().GetByUser(gomock.Eq("1")).Return(checkinsoutsData, nil).Times(1)
			},
		},
		{
			name:   "error",
			idUser: "2",
			mockCheckinsoutsRepo: func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface) {
				checkinsouts.EXPECT().GetByUser(gomock.Eq("2")).Return(checkinsoutsData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockCheckinsoutsRepo.NewMockCheckInOutRepoInterface(ctrl)
			tc.mockCheckinsoutsRepo(mockRepo)

			s := NewCheckInOutService(mockRepo)
			_, _ = s.GetByUser(tc.idUser)
		})
	}
}

func TestCheckRequest(t *testing.T) {
	checkinoutData := dummyCheckinout(t)

	testCases := []struct {
		name                 string
		idAttendance         string
		mockCheckinsoutsRepo func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface)
	}{
		{
			name:         "success",
			idAttendance: "1",
			mockCheckinsoutsRepo: func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface) {
				checkinsouts.EXPECT().CheckRequest(gomock.Eq("1")).Return(checkinoutData, nil).Times(1)
			},
		},
		{
			name:         "error",
			idAttendance: "2",
			mockCheckinsoutsRepo: func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface) {
				checkinsouts.EXPECT().CheckRequest(gomock.Eq("2")).Return(checkinoutData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockCheckinsoutsRepo.NewMockCheckInOutRepoInterface(ctrl)
			tc.mockCheckinsoutsRepo(mockRepo)

			s := NewCheckInOutService(mockRepo)
			_, _ = s.CheckRequest(tc.idAttendance)
		})
	}
}

func TestCheckData(t *testing.T) {
	testCases := []struct {
		name                 string
		idUser               string
		idAttendance         string
		mockCheckinsoutsRepo func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface)
	}{
		{
			name:         "success",
			idUser:       "1",
			idAttendance: "1",
			mockCheckinsoutsRepo: func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface) {
				checkinsouts.EXPECT().CheckData(gomock.Eq("1"), gomock.Eq("1")).Return(1).Times(1)
			},
		},
		{
			name:         "error",
			idUser:       "2",
			idAttendance: "2",
			mockCheckinsoutsRepo: func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface) {
				checkinsouts.EXPECT().CheckData(gomock.Eq("2"), gomock.Eq("2")).Return(0).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockCheckinsoutsRepo.NewMockCheckInOutRepoInterface(ctrl)
			tc.mockCheckinsoutsRepo(mockRepo)

			s := NewCheckInOutService(mockRepo)
			_ = s.CheckData(tc.idUser, tc.idAttendance)
		})
	}
}

func TestCheckIn(t *testing.T) {
	checkinoutData := dummyCheckinout(t)

	testCases := []struct {
		name                 string
		body                 _checkInsOutsRequest.CheckInsRequest
		mockCheckinsoutsRepo func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface)
	}{
		{
			name: "success",
			body: _checkInsOutsRequest.CheckInsRequest{
				ID:          _utils.RandomString(8),
				Temprature:  35.0,
				IsCheckIns:  true,
				IsCheckOuts: false,
			},
			mockCheckinsoutsRepo: func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface) {
				checkinsouts.EXPECT().CheckIn(gomock.Any()).Return(checkinoutData, nil).Times(1)
			},
		},
		{
			name: "error",
			body: _checkInsOutsRequest.CheckInsRequest{
				ID:          _utils.RandomString(8),
				Temprature:  35.0,
				IsCheckIns:  true,
				IsCheckOuts: false,
			},
			mockCheckinsoutsRepo: func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface) {
				checkinsouts.EXPECT().CheckIn(gomock.Any()).Return(checkinoutData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockCheckinsoutsRepo.NewMockCheckInOutRepoInterface(ctrl)
			tc.mockCheckinsoutsRepo(mockRepo)

			s := NewCheckInOutService(mockRepo)
			_, _ = s.Checkin(tc.body)
		})
	}
}

func TestCheckOut(t *testing.T) {
	checkinoutData := dummyCheckinout(t)

	testCases := []struct {
		name                 string
		idUser               string
		body                 _checkInsOutsRequest.CheckOutsRequest
		mockCheckinsoutsRepo func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface)
	}{
		{
			name:   "success",
			idUser: "1",
			body: _checkInsOutsRequest.CheckOutsRequest{
				ID:           _utils.RandomString(8),
				AttendanceID: _utils.RandomString(8),
				IsCheckOuts:  true,
			},

			mockCheckinsoutsRepo: func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface) {
				checkinsouts.EXPECT().CheckOut(gomock.Eq("1"), gomock.Any()).Return(checkinoutData, nil).Times(1)
			},
		},
		{
			name:   "error",
			idUser: "2",
			body: _checkInsOutsRequest.CheckOutsRequest{
				ID:           _utils.RandomString(8),
				AttendanceID: _utils.RandomString(8),
				IsCheckOuts:  true,
			},

			mockCheckinsoutsRepo: func(checkinsouts *_mockCheckinsoutsRepo.MockCheckInOutRepoInterface) {
				checkinsouts.EXPECT().CheckOut(gomock.Eq("2"), gomock.Any()).Return(checkinoutData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockCheckinsoutsRepo.NewMockCheckInOutRepoInterface(ctrl)
			tc.mockCheckinsoutsRepo(mockRepo)

			s := NewCheckInOutService(mockRepo)
			_, _ = s.Checkout(tc.idUser, tc.body)
		})
	}
}

func dummyCheckinsouts(t *testing.T) (checkinsouts []_checkinsoutsEntities.Checkin) {
	checkinsouts = []_checkinsoutsEntities.Checkin{
		{
			ID:          _utils.RandomString(8),
			Attendance:  _attendanceEntities.Attendance{ID: _utils.RandomString(8)},
			Temprature:  35.7,
			IsCheckIns:  true,
			IsCheckOuts: false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	return
}

func dummyCheckinout(t *testing.T) (checkinsouts _checkinsoutsEntities.Checkin) {
	checkinsouts = _checkinsoutsEntities.Checkin{
		ID:          _utils.RandomString(8),
		Attendance:  _attendanceEntities.Attendance{ID: _utils.RandomString(8)},
		Temprature:  35.0,
		IsCheckIns:  true,
		IsCheckOuts: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return
}
