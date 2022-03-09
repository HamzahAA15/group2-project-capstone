package officeService

import (
	"sirclo/project-capstone/entities/officeEntities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOffices(t *testing.T) {
	officeServiceMock := NewOfficeService(mockOfficeRepo{})
	offices, err := officeServiceMock.GetOffices()
	result, _ := []officeEntities.Office{
		{
			ID:   "asdasdas",
			Name: "asdasdsa",
		},
		{
			ID:   "lkkhjkhjg",
			Name: "bmnmb",
		},
	}, err
	assert.Nil(t, err)
	assert.Equal(t, result, offices)
}

type mockOfficeRepo struct {
}

func (m mockOfficeRepo) GetOffices() ([]officeEntities.Office, error) {
	return []officeEntities.Office{
		{
			ID:   "asdasdas",
			Name: "asdasdsa",
		},
		{
			ID:   "lkkhjkhjg",
			Name: "bmnmb",
		},
	}, nil
}
