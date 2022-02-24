package checkInOutRepository

import "sirclo/project-capstone/entities/checkinEntities"

type CheckInOutRepoInterface interface {
	Gets() ([]checkinEntities.Checkin, error)
	GetByUser(userID string) ([]checkinEntities.Checkin, error)
	CheckRequest(userID string, attendanceID string) int
	CheckData(userID string, attendanceID string) int
	CheckIn(checkinout checkinEntities.Checkin) (checkinEntities.Checkin, error)
	CheckOut(userID string, checkinout checkinEntities.Checkin) (checkinEntities.Checkin, error)
}
