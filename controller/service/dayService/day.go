package dayService

import (
	"sirclo/project-capstone/entities/dayEntities"
	"sirclo/project-capstone/repository/dayRepository"
	"sirclo/project-capstone/repository/userRepository"
	"sirclo/project-capstone/utils/request/dayRequest"
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

func (ds *dayService) GetDays(office, time string) ([]dayEntities.Day, error) {
	days, err := ds.dayRepository.GetDays(office, time)
	return days, err
}

func (ds *dayService) UpdateDays(input dayRequest.DayUpdateRequest) (dayEntities.Day, error) {
	var day dayEntities.Day

	day.ID = input.ID
	day.Quota = input.Quota

	updateDay, err := ds.dayRepository.UpdateDay(day)
	return updateDay, err
}

func (ds *dayService) CheckUserRole(loginId string) string {
	currentUser, _ := ds.userRepository.GetUser(loginId)
	if currentUser.Role != "admin" {
		return currentUser.Role
	}
	return currentUser.Role
}
