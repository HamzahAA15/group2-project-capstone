package dayRepository

import "sirclo/project-capstone/entities/dayEntities"

type DayRepoInterface interface {
	GetDays(office, time string) ([]dayEntities.Day, error)
	UpdateDay(day dayEntities.Day) (dayEntities.Day, error)
}
