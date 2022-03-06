package checkInOutRepository

import "sirclo/project-capstone/entities/checkinEntities"

type CheckInOutRepoInterface interface {
	Gets() ([]checkinEntities.Checkin, error)
	CheckRequest(attendanceID string) (checkinEntities.Checkin, error)
	CheckData(userID string, attendanceID string) int
	CheckIn(checkinout checkinEntities.Checkin) (checkinEntities.Checkin, error)
	CheckOut(userID string, checkinout checkinEntities.Checkin) (checkinEntities.Checkin, error)
}
