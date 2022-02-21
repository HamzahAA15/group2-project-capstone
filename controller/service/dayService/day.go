package dayService

import (
	"sirclo/project-capstone/entities/dayEntities"
	"sirclo/project-capstone/repository/dayRepository"
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
