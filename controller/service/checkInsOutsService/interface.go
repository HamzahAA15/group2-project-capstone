package checkInsOutsService

import (
	"sirclo/project-capstone/entities/checkinEntities"
	"sirclo/project-capstone/utils/request/checkInsOutsRequest"
)

//go:generate mockgen --destination=./../../../mocks/checkinout/service/mock_service_checkinout.go -source=interface.go
type CheckinoutServiceInterface interface {
	Gets() ([]checkinEntities.Checkin, error)
	GetByUser(userID string) ([]checkinEntities.Checkin, error)
	CheckRequest(attendanceID string) (checkinEntities.Checkin, error)
	CheckData(userID string, attendanceID string) int
	Checkin(input checkInsOutsRequest.CheckInsRequest) (checkinEntities.Checkin, error)
	Checkout(userID string, input checkInsOutsRequest.CheckOutsRequest) (checkinEntities.Checkin, error)
}
