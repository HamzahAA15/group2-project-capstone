package checkInsOutsService

import (
	"sirclo/project-capstone/entities/checkinEntities"
	"sirclo/project-capstone/repository/checkInOutRepository"
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
