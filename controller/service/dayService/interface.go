package dayService

import (
	"sirclo/project-capstone/entities/dayEntities"
	"sirclo/project-capstone/utils/request/dayRequest"
)

//go:generate mockgen --destination=./../../../mocks/day/service/mock_service_day.go -source=interface.go
type DayServiceInterface interface {
	GetDays(officeID string, date string) ([]dayEntities.Day, error)
	GetDaysID(dayId string) (dayEntities.Day, error)
	UpdateDays(input dayRequest.DayUpdateRequest) (dayEntities.Day, error)
}
