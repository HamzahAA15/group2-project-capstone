package dayService

import (
	"sirclo/project-capstone/entities/dayEntities"
	"sirclo/project-capstone/repository/dayRepository"
	"sirclo/project-capstone/utils/request/dayRequest"
)

type dayService struct {
	dayRepository dayRepository.DayRepoInterface
}

func NewDayService(repo dayRepository.DayRepoInterface) DayServiceInterface {
	return &dayService{
		dayRepository: repo,
	}
}

func (ds *dayService) GetDays() ([]dayEntities.Day, error) {
	days, err := ds.dayRepository.GetDays()
	if err != nil {
		return days, err
	}
	return days, nil
}

func (ds *dayService) UpdateDays(input dayRequest.DayUpdateRequest) (dayEntities.Day, error) {
	var day dayEntities.Day

	day.ID = input.ID
	day.Quota = input.Quota

	updateDay, err := ds.dayRepository.UpdateDay(day)
	if err != nil {
		return updateDay, err
	}
	return updateDay, nil
}
