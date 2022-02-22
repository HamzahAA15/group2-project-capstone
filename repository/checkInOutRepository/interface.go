package checkInOutRepository

import "sirclo/project-capstone/entities/checkinEntities"

type CheckInOutRepoInterface interface {
	Gets() ([]checkinEntities.Checkin, error)
}
