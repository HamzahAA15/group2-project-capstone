package dayService

import (
	"errors"
	"sirclo/project-capstone/entities/dayEntities"
	"sirclo/project-capstone/entities/officeEntities"
	"sirclo/project-capstone/entities/userEntities"
	"sirclo/project-capstone/utils/request/dayRequest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDays(t *testing.T) {
	dayServiceMock := NewDayService(mockDaysRepo{}, mockUserRepository{})
	day, err := dayServiceMock.GetDays("1", "2022-02-21")
	expected := []dayEntities.Day{
		{
			ID:             "qweqw",
			OfficeId:       officeEntities.Office{Name: "Head Office"},
			Date:           time.Date(time.Now().Year(), 2, 1, 0, 0, 0, 0, time.Local),
			Quota:          100,
			TotalApproved:  9,
			RemainingQuota: 91,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, day)
}

func TestGetDayId(t *testing.T) {
	dayServiceMock := NewDayService(mockDaysRepo{}, mockUserRepository{})
	day, err := dayServiceMock.GetDaysID("1")
	expected := dayEntities.Day{
		ID:             "1",
		OfficeId:       officeEntities.Office{Name: "Head Office"},
		Date:           time.Date(time.Now().Year(), 2, 1, 0, 0, 0, 0, time.Local),
		Quota:          100,
		TotalApproved:  9,
		RemainingQuota: 91,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, day)
}

func TestUpdateDay(t *testing.T) {
	dayServiceMock := NewDayService(mockDaysRepo{}, mockUserRepository{})
	input := dayRequest.DayUpdateRequest{
		ID:    "1",
		Quota: 80,
	}
	day, err := dayServiceMock.UpdateDays(input)
	expected := dayEntities.Day{
		ID:        "1",
		Quota:     80,
		UpdatedAt: time.Now(),
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, day)
}

type mockDaysRepo struct{}

func (m mockDaysRepo) GetDays(office_id string, date string) ([]dayEntities.Day, error) {
	return []dayEntities.Day{
		{
			ID:             "qweqw",
			OfficeId:       officeEntities.Office{Name: "Head Office"},
			Date:           time.Date(time.Now().Year(), 2, 1, 0, 0, 0, 0, time.Local),
			Quota:          100,
			TotalApproved:  9,
			RemainingQuota: 91,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}, nil
}

func (m mockDaysRepo) GetDayID(dayID string) (dayEntities.Day, error) {
	return dayEntities.Day{
		ID:             dayID,
		OfficeId:       officeEntities.Office{Name: "Head Office"},
		Date:           time.Date(time.Now().Year(), 2, 1, 0, 0, 0, 0, time.Local),
		Quota:          100,
		TotalApproved:  9,
		RemainingQuota: 91,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (m mockDaysRepo) UpdateDay(day dayEntities.Day) (dayEntities.Day, error) {
	return dayEntities.Day{
		ID:        "1",
		Quota:     80,
		UpdatedAt: time.Now(),
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
