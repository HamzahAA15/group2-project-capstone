package dayService

import (
	"sirclo/project-capstone/entities/dayEntities"
	"sirclo/project-capstone/utils/request/dayRequest"
)

type DayServiceInterface interface {
	GetDays() ([]dayEntities.Day, error)
	UpdateDays(input dayRequest.DayUpdateRequest) (dayEntities.Day, error)
}
