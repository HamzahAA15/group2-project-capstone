package attendanceService

import (
	"errors"
	"sirclo/project-capstone/entities/attendanceEntities"
	"sirclo/project-capstone/entities/dayEntities"
	"sirclo/project-capstone/entities/userEntities"
	"sirclo/project-capstone/utils/request/attendanceRequest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAttendancesById(t *testing.T) {
	attServiceMock := NewAttendanceService(mockAttendanceRepo{}, mockUserRepository{})
	_, name, err := attServiceMock.GetAttendancesById("3")
	assert.Nil(t, err)
	assert.Equal(t, "jundana", name)
}

func TestGetAttendancesCurrentUser(t *testing.T) {
	attServiceMock := NewAttendanceService(mockAttendanceRepo{}, mockUserRepository{})
	att, err := attServiceMock.GetAttendancesCurrentUser("1", "approved", "asc")
	expected := []attendanceEntities.Attendance{
		{
			ID:       "1",
			Day:      dayEntities.Day{Date: time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local)},
			Office:   "Head Office",
			Employee: userEntities.User{Avatar: "https://ui-avatars.com/api/?name=Hamzah", Email: "hamzah@sirclo.com", Nik: "327658403093", Name: "Hamzah"},
			Status:   "approved",
			Notes:    "",
			Admin:    userEntities.User{Name: "jundana"},
		},
		{
			ID:       "2",
			Day:      dayEntities.Day{Date: time.Date(time.Now().Year(), time.Now().Month()+1, 1, 0, 0, 0, 0, time.Local)},
			Office:   "Head Office",
			Employee: userEntities.User{Avatar: "https://ui-avatars.com/api/?name=Hamzah", Email: "hamzah@sirclo.com", Nik: "327658403093", Name: "Hamzah"},
			Status:   "approved",
			Notes:    "",
			Admin:    userEntities.User{Name: "felicia"},
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, att)
}

func TestGetAttendancesRangeDate(t *testing.T) {
	attServiceMock := NewAttendanceService(mockAttendanceRepo{}, mockUserRepository{})
	att, err := attServiceMock.GetAttendancesRangeDate("hamzah@sirclo.com", time.Now().String(), time.Now().AddDate(0, 0, 2).String(), "approved", "5ca05e7d-9536-4893-bdb6-0c04de89b047", "asc")
	expected := []attendanceEntities.Attendance{
		{
			ID:       "1",
			Day:      dayEntities.Day{Date: time.Now()},
			OfficeId: "5ca05e7d-9536-4893-bdb6-0c04de89b047",
			Office:   "Head Office",
			Employee: userEntities.User{Avatar: "https://ui-avatars.com/api/?name=Hamzah", Email: "hamzah@sirclo.com", Nik: "327658403093", Name: "Hamzah"},
			Status:   "approved",
			Notes:    "",
			Admin:    userEntities.User{Name: "jundana"},
		},
		{
			ID:       "4",
			Day:      dayEntities.Day{Date: time.Now().AddDate(0, 0, 1)},
			OfficeId: "5ca05e7d-9536-4893-bdb6-0c04de89b047",
			Office:   "Head Office",
			Employee: userEntities.User{Avatar: "https://ui-avatars.com/api/?name=Hamzah", Email: "hamzah@sirclo.com", Nik: "327658403093", Name: "hamzah"},
			Status:   "approved",
			Notes:    "",
			Admin:    userEntities.User{Name: "felicia"},
		},
		{
			ID:       "8",
			Day:      dayEntities.Day{Date: time.Now().AddDate(0, 0, 2)},
			OfficeId: "5ca05e7d-9536-4893-bdb6-0c04de89b047",
			Office:   "Head Office",
			Employee: userEntities.User{Avatar: "https://ui-avatars.com/api/?name=Hamzah", Email: "hamzah@sirclo.com", Nik: "327658403093", Name: "hamzah"},
			Status:   "approved",
			Notes:    "",
			Admin:    userEntities.User{Name: "felicia"},
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, att)
}

func TestGetAttendancesRangeDateNil(t *testing.T) {
	attServiceMock := NewAttendanceService(mockAttendanceRepo{}, mockUserRepository{})
	att, err := attServiceMock.GetAttendancesRangeDate("hamzah@sirclo.com", "", "", "approved", "5ca05e7d-9536-4893-bdb6-0c04de89b047", "asc")
	expected := []attendanceEntities.Attendance{
		{
			ID:       "1",
			Day:      dayEntities.Day{Date: time.Now()},
			OfficeId: "5ca05e7d-9536-4893-bdb6-0c04de89b047",
			Office:   "Head Office",
			Employee: userEntities.User{Avatar: "https://ui-avatars.com/api/?name=Hamzah", Email: "hamzah@sirclo.com", Nik: "327658403093", Name: "Hamzah"},
			Status:   "approved",
			Notes:    "",
			Admin:    userEntities.User{Name: "jundana"},
		},
		{
			ID:       "4",
			Day:      dayEntities.Day{Date: time.Now().AddDate(0, 0, 1)},
			OfficeId: "5ca05e7d-9536-4893-bdb6-0c04de89b047",
			Office:   "Head Office",
			Employee: userEntities.User{Avatar: "https://ui-avatars.com/api/?name=Hamzah", Email: "hamzah@sirclo.com", Nik: "327658403093", Name: "hamzah"},
			Status:   "approved",
			Notes:    "",
			Admin:    userEntities.User{Name: "felicia"},
		},
		{
			ID:       "8",
			Day:      dayEntities.Day{Date: time.Now().AddDate(0, 0, 2)},
			OfficeId: "5ca05e7d-9536-4893-bdb6-0c04de89b047",
			Office:   "Head Office",
			Employee: userEntities.User{Avatar: "https://ui-avatars.com/api/?name=Hamzah", Email: "hamzah@sirclo.com", Nik: "327658403093", Name: "hamzah"},
			Status:   "approved",
			Notes:    "",
			Admin:    userEntities.User{Name: "felicia"},
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, att)
}

func TestIsCheckins(t *testing.T) {
	attServiceMock := NewAttendanceService(mockAttendanceRepo{}, mockUserRepository{})
	att, err := attServiceMock.IsCheckins()
	expected := []attendanceEntities.Attendance{
		{
			ID:            "2",
			Employee:      userEntities.User{Email: "hamzah@sirclo.com", Name: "hamzah", Avatar: "https://ui-avatars.com/api/?name=Hamzah"},
			Office:        "Head Office",
			StatusCheckin: "Checkin",
		},
		{
			ID:            "5",
			Employee:      userEntities.User{Email: "lukman@sirclo.com", Name: "lukman", Avatar: "https://ui-avatars.com/api/?name=Lukman"},
			Office:        "Head Office",
			StatusCheckin: "Checkin",
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, att)
}

func TestCreateAttendance(t *testing.T) {
	attServiceMock := NewAttendanceService(mockAttendanceRepo{}, mockUserRepository{})
	att1 := attendanceRequest.CreateAttRequest{
		Day: "1",
	}
	att, err := attServiceMock.CreateAttendance("12", att1)
	expected := attendanceEntities.Attendance{
		ID:       "1",
		Day:      dayEntities.Day{ID: "1"},
		Employee: userEntities.User{ID: "12"},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, att)
}
func TestUpdateAttendance(t *testing.T) {
	attServiceMock := NewAttendanceService(mockAttendanceRepo{}, mockUserRepository{})
	att1 := attendanceRequest.UpdateAttRequest{
		ID:     "1",
		Status: "approved",
		Notes:  "",
		Admin:  "1",
	}
	att, err := attServiceMock.UpdateAttendance("1", att1)
	expected := attendanceEntities.Attendance{
		ID:     "1",
		Status: "approved",
		Notes:  "",
		Admin:  userEntities.User{ID: "1"},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, att)
}

type mockAttendanceRepo struct{}

func (m mockAttendanceRepo) GetAttendancesById(attID string) (string, string, error) {
	return "1", "jundana", nil
}

func (m mockAttendanceRepo) GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, officeId, order string) ([]attendanceEntities.Attendance, error) {
	return []attendanceEntities.Attendance{
		{
			ID:       "1",
			Day:      dayEntities.Day{Date: time.Now()},
			OfficeId: officeId,
			Office:   "Head Office",
			Employee: userEntities.User{Avatar: "https://ui-avatars.com/api/?name=Hamzah", Email: employeeEmail, Nik: "327658403093", Name: "Hamzah"},
			Status:   status,
			Notes:    "",
			Admin:    userEntities.User{Name: "jundana"},
		},
		{
			ID:       "4",
			Day:      dayEntities.Day{Date: time.Now().AddDate(0, 0, 1)},
			OfficeId: officeId,
			Office:   "Head Office",
			Employee: userEntities.User{Avatar: "https://ui-avatars.com/api/?name=Hamzah", Email: employeeEmail, Nik: "327658403093", Name: "hamzah"},
			Status:   status,
			Notes:    "",
			Admin:    userEntities.User{Name: "felicia"},
		},
		{
			ID:       "8",
			Day:      dayEntities.Day{Date: time.Now().AddDate(0, 0, 2)},
			OfficeId: officeId,
			Office:   "Head Office",
			Employee: userEntities.User{Avatar: "https://ui-avatars.com/api/?name=Hamzah", Email: employeeEmail, Nik: "327658403093", Name: "hamzah"},
			Status:   status,
			Notes:    "",
			Admin:    userEntities.User{Name: "felicia"},
		},
	}, nil
}

func (m mockAttendanceRepo) GetAttendancesCurrentUser(userId, status, order string) ([]attendanceEntities.Attendance, error) {
	return []attendanceEntities.Attendance{
		{
			ID:       "1",
			Day:      dayEntities.Day{Date: time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local)},
			Office:   "Head Office",
			Employee: userEntities.User{Avatar: "https://ui-avatars.com/api/?name=Hamzah", Email: "hamzah@sirclo.com", Nik: "327658403093", Name: "Hamzah"},
			Status:   status,
			Notes:    "",
			Admin:    userEntities.User{Name: "jundana"},
		},
		{
			ID:       "2",
			Day:      dayEntities.Day{Date: time.Date(time.Now().Year(), time.Now().Month()+1, 1, 0, 0, 0, 0, time.Local)},
			Office:   "Head Office",
			Employee: userEntities.User{Avatar: "https://ui-avatars.com/api/?name=Hamzah", Email: "hamzah@sirclo.com", Nik: "327658403093", Name: "Hamzah"},
			Status:   status,
			Notes:    "",
			Admin:    userEntities.User{Name: "felicia"},
		},
	}, nil
}

func (m mockAttendanceRepo) IsCheckins() ([]attendanceEntities.Attendance, error) {
	return []attendanceEntities.Attendance{
		{
			ID:            "2",
			Employee:      userEntities.User{Email: "hamzah@sirclo.com", Name: "hamzah", Avatar: "https://ui-avatars.com/api/?name=Hamzah"},
			Office:        "Head Office",
			StatusCheckin: "Checkin",
		},
		{
			ID:            "5",
			Employee:      userEntities.User{Email: "lukman@sirclo.com", Name: "lukman", Avatar: "https://ui-avatars.com/api/?name=Lukman"},
			Office:        "Head Office",
			StatusCheckin: "Checkin",
		},
	}, nil
}

func (m mockAttendanceRepo) CreateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error) {
	return attendanceEntities.Attendance{
		ID:       "1",
		Day:      dayEntities.Day{ID: att.Day.ID},
		Employee: userEntities.User{ID: att.Employee.ID},
	}, nil
}

func (m mockAttendanceRepo) UpdateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error) {
	return attendanceEntities.Attendance{
		ID:     att.ID,
		Status: att.Status,
		Notes:  att.Notes,
		Admin:  userEntities.User{ID: att.Admin.ID},
	}, nil
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
	return userEntities.User{}, nil
}

func (m mockUserRepository) Login(identity string) (userEntities.User, error) {
	return userEntities.User{}, nil
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
