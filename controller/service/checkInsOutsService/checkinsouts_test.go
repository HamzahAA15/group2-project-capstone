package checkInsOutsService

import (
	"sirclo/project-capstone/entities/attendanceEntities"
	"sirclo/project-capstone/entities/checkinEntities"
	"sirclo/project-capstone/entities/userEntities"
	"sirclo/project-capstone/utils/request/checkInsOutsRequest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGets(t *testing.T) {
	cekServiceMock := NewCheckInOutService(mockCheckInsOutRepo{})
	cek, err := cekServiceMock.Gets()
	expected := []checkinEntities.Checkin{
		{
			ID:         "1",
			Attendance: attendanceEntities.Attendance{ID: "1"},
			Temprature: 35.7,
			IsCheckIns: true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         "2",
			Attendance: attendanceEntities.Attendance{ID: "2"},
			Temprature: 36.7,
			IsCheckIns: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, cek)
}

func TestGetByUser(t *testing.T) {
	cekServiceMock := NewCheckInOutService(mockCheckInsOutRepo{})
	cek, err := cekServiceMock.GetByUser("3")
	expected := []checkinEntities.Checkin{
		{
			ID:         "1",
			Attendance: attendanceEntities.Attendance{ID: "1", Employee: userEntities.User{ID: "3"}},
			Temprature: 35.5,
			IsCheckIns: true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         "2",
			Attendance: attendanceEntities.Attendance{ID: "2", Employee: userEntities.User{ID: "3"}},
			Temprature: 36.4,
			IsCheckIns: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, cek)
}

func TestCheckRequest(t *testing.T) {
	cekServiceMock := NewCheckInOutService(mockCheckInsOutRepo{})
	cek, err := cekServiceMock.CheckRequest("7")
	expected := checkinEntities.Checkin{
		ID:         "1",
		Attendance: attendanceEntities.Attendance{ID: "7", Employee: userEntities.User{ID: "3"}, Status: "approved"},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, cek)
}

func TestCheckData(t *testing.T) {
	cekServiceMock := NewCheckInOutService(mockCheckInsOutRepo{})
	cek := cekServiceMock.CheckData("3", "2")
	assert.Equal(t, 1, cek)
}

func TestCheckIn(t *testing.T) {
	cekServiceMock := NewCheckInOutService(mockCheckInsOutRepo{})
	input := checkInsOutsRequest.CheckInsRequest{
		ID:           "1",
		AttendanceID: "2",
		Temprature:   35.2,
		IsCheckIns:   true,
		IsCheckOuts:  false,
	}
	cek, err := cekServiceMock.Checkin(input)
	expected := checkinEntities.Checkin{
		ID:         "1",
		Attendance: attendanceEntities.Attendance{ID: "2"}, Temprature: 35.2,
		IsCheckIns:  true,
		IsCheckOuts: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, cek)
}

func TestCheckOut(t *testing.T) {
	cekServiceMock := NewCheckInOutService(mockCheckInsOutRepo{})
	input := checkInsOutsRequest.CheckOutsRequest{
		ID:           "1",
		AttendanceID: "2",
		IsCheckOuts:  false,
	}
	cek, err := cekServiceMock.Checkout("3", input)
	expected := checkinEntities.Checkin{
		ID:          "1",
		Attendance:  attendanceEntities.Attendance{ID: "2", Employee: userEntities.User{ID: "3"}},
		IsCheckOuts: true,
		UpdatedAt:   time.Now(),
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, cek)
}

type mockCheckInsOutRepo struct{}

func (m mockCheckInsOutRepo) Gets() ([]checkinEntities.Checkin, error) {
	return []checkinEntities.Checkin{
		{
			ID:         "1",
			Attendance: attendanceEntities.Attendance{ID: "1"},
			Temprature: 35.7,
			IsCheckIns: true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         "2",
			Attendance: attendanceEntities.Attendance{ID: "2"},
			Temprature: 36.7,
			IsCheckIns: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}, nil
}

func (m mockCheckInsOutRepo) GetByUser(userID string) ([]checkinEntities.Checkin, error) {
	return []checkinEntities.Checkin{
		{
			ID:         "1",
			Attendance: attendanceEntities.Attendance{ID: "1", Employee: userEntities.User{ID: userID}},
			Temprature: 35.5,
			IsCheckIns: true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         "2",
			Attendance: attendanceEntities.Attendance{ID: "2", Employee: userEntities.User{ID: userID}},
			Temprature: 36.4,
			IsCheckIns: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}, nil
}

func (m mockCheckInsOutRepo) CheckRequest(attendanceID string) (checkinEntities.Checkin, error) {
	return checkinEntities.Checkin{
		ID:         "1",
		Attendance: attendanceEntities.Attendance{ID: attendanceID, Employee: userEntities.User{ID: "3"}, Status: "approved"},
	}, nil
}

func (m mockCheckInsOutRepo) CheckData(userID, attendanceID string) int {
	return 1
}

func (m mockCheckInsOutRepo) CheckIn(checkin checkinEntities.Checkin) (checkinEntities.Checkin, error) {
	return checkinEntities.Checkin{
		ID:          "1",
		Attendance:  attendanceEntities.Attendance{ID: "2"},
		Temprature:  35.2,
		IsCheckIns:  true,
		IsCheckOuts: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m mockCheckInsOutRepo) CheckOut(userID string, checkin checkinEntities.Checkin) (checkinEntities.Checkin, error) {
	return checkinEntities.Checkin{
		ID:          "1",
		Attendance:  attendanceEntities.Attendance{ID: "2", Employee: userEntities.User{ID: userID}},
		IsCheckOuts: true,
		UpdatedAt:   time.Now(),
	}, nil
}
