package checkInsOutsService

import (
	"sirclo/project-capstone/entities/checkinEntities"
	"sirclo/project-capstone/utils/request/checkInsOutsRequest"
)

type CheckinoutServiceInterface interface {
	Gets() ([]checkinEntities.Checkin, error)
	GetByUser(userID string) ([]checkinEntities.Checkin, error)
	Checkin(input checkInsOutsRequest.CheckInsRequest) (checkinEntities.Checkin, error)
}
