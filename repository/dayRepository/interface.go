package dayRepository

import "sirclo/project-capstone/entities/dayEntities"

type DayRepoInterface interface {
	GetDays() ([]dayEntities.Day, error)
	UpdateDay(day dayEntities.Day) (dayEntities.Day, error)
}
