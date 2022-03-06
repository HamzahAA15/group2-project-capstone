package dayService

import (
	"sirclo/project-capstone/entities/dayEntities"
	"sirclo/project-capstone/repository/dayRepository"
	"sirclo/project-capstone/repository/userRepository"
	"sirclo/project-capstone/utils/request/dayRequest"
	"time"
)

type dayService struct {
	dayRepository  dayRepository.DayRepoInterface
	userRepository userRepository.UserRepoInterface
}

func NewDayService(dayRepo dayRepository.DayRepoInterface, userRepo userRepository.UserRepoInterface) DayServiceInterface {
	return &dayService{
		dayRepository:  dayRepo,
		userRepository: userRepo,
	}
}

func (ds *dayService) GetDays(officeID string, date string) ([]dayEntities.Day, error) {
	days, err := ds.dayRepository.GetDays(officeID, date)
	return days, err
}

func (ds *dayService) GetDaysID(dayId string) (dayEntities.Day, error) {
	days, err := ds.dayRepository.GetDayID(dayId)
	return days, err
}

func (ds *dayService) UpdateDays(input dayRequest.DayUpdateRequest) (dayEntities.Day, error) {
	var day dayEntities.Day

	day.ID = input.ID
	day.Quota = input.Quota
	day.UpdatedAt = time.Now()

	updateDay, err := ds.dayRepository.UpdateDay(day)
	return updateDay, err
}
