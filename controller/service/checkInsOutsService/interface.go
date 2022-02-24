package checkInsOutsService

import (
	"sirclo/project-capstone/entities/checkinEntities"
	"sirclo/project-capstone/utils/request/checkInsOutsRequest"
)

type CheckinoutServiceInterface interface {
	Gets() ([]checkinEntities.Checkin, error)
	GetByUser(userID string) ([]checkinEntities.Checkin, error)
	CheckRequest(userID string, attendanceID string) int
	CheckData(userID string, attendanceID string) int
	Checkin(input checkInsOutsRequest.CheckInsRequest) (checkinEntities.Checkin, error)
	Checkout(userID string, input checkInsOutsRequest.CheckOutsRequest) (checkinEntities.Checkin, error)
}
