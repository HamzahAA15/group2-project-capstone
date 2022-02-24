package checkInsOutsService

import (
	"sirclo/project-capstone/entities/checkinEntities"
	"sirclo/project-capstone/repository/checkInOutRepository"
	"sirclo/project-capstone/utils/request/checkInsOutsRequest"
	"time"

	"github.com/google/uuid"
)

type checkinoutService struct {
	checkinoutRepository checkInOutRepository.CheckInOutRepoInterface
}

func NewCheckInOutService(repo checkInOutRepository.CheckInOutRepoInterface) CheckinoutServiceInterface {
	return &checkinoutService{
		checkinoutRepository: repo,
	}
}

func (cs *checkinoutService) Gets() ([]checkinEntities.Checkin, error) {
	checkinsout, err := cs.checkinoutRepository.Gets()
	return checkinsout, err
}

func (cs *checkinoutService) GetByUser(userID string) ([]checkinEntities.Checkin, error) {
	checkinsout, err := cs.checkinoutRepository.GetByUser(userID)
	return checkinsout, err
}

func (cs *checkinoutService) CheckRequest(userID string, attendanceID string) int {
	countData := cs.checkinoutRepository.CheckRequest(userID, attendanceID)
	return countData
}

func (cs *checkinoutService) CheckData(userID string, attendanceID string) int {
	countData := cs.checkinoutRepository.CheckData(userID, attendanceID)
	return countData
}

func (cs *checkinoutService) Checkin(input checkInsOutsRequest.CheckInsRequest) (checkinEntities.Checkin, error) {
	checkins := checkinEntities.Checkin{}
	checkins.ID = uuid.New().String()
	checkins.AttendanceID = input.AttendanceID
	checkins.Temprature = input.Temprature
	checkins.IsCheckIns = true
	checkins.IsCheckOuts = false
	checkins.CreatedAt = time.Now()
	checkins.UpdatedAt = time.Now()

	createCheckins, err := cs.checkinoutRepository.CheckIn(checkins)
	return createCheckins, err
}

func (cs *checkinoutService) Checkout(userID string, input checkInsOutsRequest.CheckOutsRequest) (checkinEntities.Checkin, error) {
	checkouts := checkinEntities.Checkin{}
	checkouts.ID = input.ID
	checkouts.AttendanceID = input.AttendanceID
	checkouts.IsCheckOuts = true
	checkouts.UpdatedAt = time.Now()

	checkOuts, err := cs.checkinoutRepository.CheckOut(userID, checkouts)
	return checkOuts, err
}
