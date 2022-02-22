package checkInsOutsService

import (
	"sirclo/project-capstone/entities/checkinEntities"
)

type CheckinoutServiceInterface interface {
	Gets() ([]checkinEntities.Checkin, error)
}
