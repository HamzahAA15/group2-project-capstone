package officeService

import (
	"database/sql"
	"testing"

	_officeEntities "sirclo/project-capstone/entities/officeEntities"
	_mockRepo "sirclo/project-capstone/mocks/office/repository"

	gomock "github.com/golang/mock/gomock"
)

func TestGetOffices(t *testing.T) {
	testCases := []struct {
		name           string
		mockOfficeRepo func(office *_mockRepo.MockOfficeRepoInterface)
	}{
		{
			name: "success",
			mockOfficeRepo: func(office *_mockRepo.MockOfficeRepoInterface) {
				office.EXPECT().GetOffices().Return([]_officeEntities.Office{}, nil).Times(1)
			},
		},
		{
			name: "error",
			mockOfficeRepo: func(office *_mockRepo.MockOfficeRepoInterface) {
				office.EXPECT().GetOffices().Return([]_officeEntities.Office{}, sql.ErrConnDone).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockRepo.NewMockOfficeRepoInterface(ctrl)
			tc.mockOfficeRepo(mockRepo)

			s := NewOfficeService(mockRepo)
			_, _ = s.GetOffices()
		})
	}
}
