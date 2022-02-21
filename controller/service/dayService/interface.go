package dayService

import "sirclo/project-capstone/entities/dayEntities"

type DayServiceInterface interface {
	GetDays() ([]dayEntities.Day, error)
}
