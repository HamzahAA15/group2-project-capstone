package dayRepository

import (
	"sirclo/project-capstone/entities/dayEntities"
)

type DayRepoInterface interface {
	GetDays(office_id string, date string) ([]dayEntities.Day, error)
	UpdateDay(day dayEntities.Day) (dayEntities.Day, error)
}
