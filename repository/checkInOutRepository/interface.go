package checkInOutRepository

import "sirclo/project-capstone/entities/checkinEntities"

//go:generate mockgen --destination=./../../mocks/checkinout/repository/mock_repository_checkinout.go -source=interface.go
type CheckInOutRepoInterface interface {
	Gets() ([]checkinEntities.Checkin, error)
	GetByUser(userID string) ([]checkinEntities.Checkin, error)
	CheckRequest(attendanceID string) (checkinEntities.Checkin, error)
	CheckData(userID string, attendanceID string) int
	CheckIn(checkinout checkinEntities.Checkin) (checkinEntities.Checkin, error)
	CheckOut(userID string, checkinout checkinEntities.Checkin) (checkinEntities.Checkin, error)
}
