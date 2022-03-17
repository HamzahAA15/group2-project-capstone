package attendanceService

import (
	"errors"
	"testing"
	"time"

	_attendanceEntities "sirclo/project-capstone/entities/attendanceEntities"
	_dayEntities "sirclo/project-capstone/entities/dayEntities"
	_userEntities "sirclo/project-capstone/entities/userEntities"
	_mockAttendanceRepo "sirclo/project-capstone/mocks/attendance/repository"
	_mockUserRepo "sirclo/project-capstone/mocks/user/repository"
	_utils "sirclo/project-capstone/utils"
	_attendanceRequest "sirclo/project-capstone/utils/request/attendanceRequest"

	gomock "github.com/golang/mock/gomock"
)

func TestGetAttendancesById(t *testing.T) {
	testCases := []struct {
		name         string
		idAttendance string
		mockRepo     func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface)
	}{
		{
			name:         "Success Get Data",
			idAttendance: "1",
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().GetAttendancesById(gomock.Eq("1")).Return(_utils.RandomString(8), _utils.RandomString(8), nil).Times(1)
			},
		},
		{
			name:         "Error Get Data",
			idAttendance: "2",
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().GetAttendancesById(gomock.Eq("2")).Return(_utils.RandomString(8), _utils.RandomString(8), errors.New("Error Get Data")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockAttendanceRepo := _mockAttendanceRepo.NewMockAttendanceRepoInterface(ctrl)
			mockUserRepo := _mockUserRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockAttendanceRepo, mockUserRepo)

			s := NewAttendanceService(mockAttendanceRepo, mockUserRepo)
			_, _, _ = s.GetAttendancesById(tc.idAttendance)
		})
	}
}

func TestGetAttendancesCurrentUser(t *testing.T) {
	attendancesData := dummyAttendances(t)

	testCases := []struct {
		name     string
		idUser   string
		status   string
		order    string
		mockRepo func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface)
	}{
		{
			name:   "Success Get Data",
			idUser: "1",
			status: "approved",
			order:  "asc",
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().GetAttendancesCurrentUser(gomock.Eq("1"), gomock.Eq("approved"), gomock.Eq("asc")).Return(attendancesData, nil).Times(1)
			},
		},
		{
			name:   "Error Get Data",
			idUser: "2",
			status: "rejected",
			order:  "desc",
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().GetAttendancesCurrentUser(gomock.Eq("2"), gomock.Eq("rejected"), gomock.Eq("desc")).Return(attendancesData, errors.New("Error Get Data")).Times(1)
			},
		}}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockAttendanceRepo := _mockAttendanceRepo.NewMockAttendanceRepoInterface(ctrl)
			mockUserRepo := _mockUserRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockAttendanceRepo, mockUserRepo)

			s := NewAttendanceService(mockAttendanceRepo, mockUserRepo)
			_, _ = s.GetAttendancesCurrentUser(tc.idUser, tc.status, tc.order)
		})
	}
}

func TestGetAttendancesRangeDate(t *testing.T) {
	attendancesData := dummyAttendances(t)

	testCases := []struct {
		name          string
		employeeEmail string
		dateStart     string
		dateEnd       string
		status        string
		idOffice      string
		orderBy       string
		mockRepo      func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface)
	}{
		{
			name:          "Success Get Data",
			employeeEmail: "test@mail.com",
			dateStart:     "2001-01-01",
			dateEnd:       "2001-01-01",
			status:        "approved",
			idOffice:      "5ca05e7d-9536-4893-bdb6-0c04de89b047",
			orderBy:       "asc",
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().GetAttendancesRangeDate(gomock.Eq("test@mail.com"), gomock.Eq("2001-01-01"), gomock.Eq("2001-01-01"), gomock.Eq("approved"), gomock.Eq("5ca05e7d-9536-4893-bdb6-0c04de89b047"), gomock.Eq("asc")).Return(attendancesData, nil).Times(1)
			},
		},
		{
			name:          "Error Get Data",
			employeeEmail: "fail@mail.com",
			dateStart:     "2002-02-02",
			dateEnd:       "2002-02-02",
			status:        "rejected",
			idOffice:      "5ca05e7d-9536-4893-bdb6-0c04de89b047",
			orderBy:       "desc",
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().GetAttendancesRangeDate(gomock.Eq("fail@mail.com"), gomock.Eq("2002-02-02"), gomock.Eq("2002-02-02"), gomock.Eq("rejected"), gomock.Eq("5ca05e7d-9536-4893-bdb6-0c04de89b047"), gomock.Eq("desc")).Return(attendancesData, errors.New("Error Get Data")).Times(1)
			},
		},
		{
			name:          "Success Get Data if DateStart is Empty",
			employeeEmail: "test@mail.com",
			dateStart:     "",
			dateEnd:       "2003-03-03",
			status:        "approved",
			idOffice:      "5ca05e7d-9536-4893-bdb6-0c04de89b047",
			orderBy:       "asc",
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().GetAttendancesRangeDate(gomock.Eq("test@mail.com"), gomock.Any(), gomock.Eq("2003-03-03"), gomock.Eq("approved"), gomock.Eq("5ca05e7d-9536-4893-bdb6-0c04de89b047"), gomock.Eq("asc")).Return([]_attendanceEntities.Attendance{}, nil).Times(1)
			},
		},
		{
			name:          "Success Get Data if DateEnd is Empty",
			employeeEmail: "test@mail.com",
			dateStart:     "2004-04-04",
			dateEnd:       "",
			status:        "approved",
			idOffice:      "5ca05e7d-9536-4893-bdb6-0c04de89b047",
			orderBy:       "asc",
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().GetAttendancesRangeDate(gomock.Eq("test@mail.com"), gomock.Eq("2004-04-04"), gomock.Any(), gomock.Eq("approved"), gomock.Eq("5ca05e7d-9536-4893-bdb6-0c04de89b047"), gomock.Eq("asc")).Return(attendancesData, nil).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockAttendanceRepo := _mockAttendanceRepo.NewMockAttendanceRepoInterface(ctrl)
			mockUserRepo := _mockUserRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockAttendanceRepo, mockUserRepo)

			s := NewAttendanceService(mockAttendanceRepo, mockUserRepo)
			_, _ = s.GetAttendancesRangeDate(tc.employeeEmail, tc.dateStart, tc.dateEnd, tc.status, tc.idOffice, tc.orderBy)
		})
	}
}

func TestIsCheckins(t *testing.T) {
	attendancesData := dummyAttendances(t)

	testCases := []struct {
		name     string
		mockRepo func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface)
	}{
		{
			name: "Success Data",
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().IsCheckins().Return(attendancesData, nil).Times(1)
			},
		},
		{
			name: "Error Data",
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().IsCheckins().Return(attendancesData, errors.New("Error Get Data")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockAttendanceRepo := _mockAttendanceRepo.NewMockAttendanceRepoInterface(ctrl)
			mockUserRepo := _mockUserRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockAttendanceRepo, mockUserRepo)

			s := NewAttendanceService(mockAttendanceRepo, mockUserRepo)
			_, _ = s.IsCheckins()
		})
	}
}

func TestCreateAttendance(t *testing.T) {
	attendanceData := dummyAttendance(t)

	testCases := []struct {
		name     string
		idLogin  string
		body     _attendanceRequest.CreateAttRequest
		mockRepo func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface)
	}{
		{
			name:    "Success Create",
			idLogin: "1",
			body: _attendanceRequest.CreateAttRequest{
				Day: "1",
			},
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().CreateAttendance(gomock.Any()).Return(attendanceData, nil).Times(1)
			},
		},
		{
			name:    "Error Create",
			idLogin: "2",
			body: _attendanceRequest.CreateAttRequest{
				Day: "2",
			},
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().CreateAttendance(gomock.Any()).Return(attendanceData, errors.New("Error Create")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockAttendanceRepo := _mockAttendanceRepo.NewMockAttendanceRepoInterface(ctrl)
			mockUserRepo := _mockUserRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockAttendanceRepo, mockUserRepo)

			s := NewAttendanceService(mockAttendanceRepo, mockUserRepo)
			_, _ = s.CreateAttendance(tc.idLogin, tc.body)
		})
	}
}

func TestUpdateAttendance(t *testing.T) {
	attendanceData := dummyAttendance(t)

	testCases := []struct {
		name     string
		idLogin  string
		body     _attendanceRequest.UpdateAttRequest
		mockRepo func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface)
	}{
		{
			name:    "Success Update",
			idLogin: "1",
			body: _attendanceRequest.UpdateAttRequest{
				ID:     _utils.RandomString(8),
				Status: "approved",
				Notes:  _utils.RandomString(20),
				Admin:  _utils.RandomString(8),
			},
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().UpdateAttendance(gomock.Any()).Return(attendanceData, nil).Times(1)
			},
		},
		{
			name:    "Error Update",
			idLogin: "2",
			body: _attendanceRequest.UpdateAttRequest{
				ID:     _utils.RandomString(8),
				Status: "approved",
				Notes:  _utils.RandomString(20),
				Admin:  _utils.RandomString(8),
			},
			mockRepo: func(attendance *_mockAttendanceRepo.MockAttendanceRepoInterface, user *_mockUserRepo.MockUserRepoInterface) {
				attendance.EXPECT().UpdateAttendance(gomock.Any()).Return(attendanceData, errors.New("Error Update")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockAttendanceRepo := _mockAttendanceRepo.NewMockAttendanceRepoInterface(ctrl)
			mockUserRepo := _mockUserRepo.NewMockUserRepoInterface(ctrl)
			tc.mockRepo(mockAttendanceRepo, mockUserRepo)

			s := NewAttendanceService(mockAttendanceRepo, mockUserRepo)
			_, _ = s.UpdateAttendance(tc.idLogin, tc.body)
		})
	}
}

func dummyAttendances(t *testing.T) (attendances []_attendanceEntities.Attendance) {
	attendances = []_attendanceEntities.Attendance{
		{
			ID: _utils.RandomString(8),
			Day: _dayEntities.Day{
				Date: time.Now(),
			},
			OfficeId: _utils.RandomString(8),
			Office:   "Head Office",
			Employee: _userEntities.User{
				Avatar: "https://ui-avatars.com/api/?name=" + _utils.RandomString(5),
				Email:  _utils.RandomString(5) + "@" + _utils.RandomString(4) + ".com",
				Nik:    "327658403093",
				Name:   _utils.RandomString(5),
			},
			Status: "approved",
			Notes:  _utils.RandomString(20),
			Admin: _userEntities.User{
				Name: _utils.RandomString(5),
			},
			StatusCheckin: "checkin",
		},
	}

	return
}

func dummyAttendance(t *testing.T) (attendance _attendanceEntities.Attendance) {
	attendance = _attendanceEntities.Attendance{
		ID: _utils.RandomString(8),
		Day: _dayEntities.Day{
			Date: time.Now(),
		},
		OfficeId: _utils.RandomString(8),
		Office:   "Head Office",
		Employee: _userEntities.User{
			Avatar: "https://ui-avatars.com/api/?name=" + _utils.RandomString(5),
			Email:  _utils.RandomString(5) + "@" + _utils.RandomString(4) + ".com",
			Nik:    "327658403093",
			Name:   _utils.RandomString(5),
		},
		Status: "approved",
		Notes:  _utils.RandomString(20),
		Admin: _userEntities.User{
			Name: _utils.RandomString(5),
		},
		StatusCheckin: "checkin",
	}

	return
}
