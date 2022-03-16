package dayRepository

import (
	"sirclo/project-capstone/entities/dayEntities"
)

//go:generate mockgen --destination=./../../mocks/day/repository/mock_repository_day.go -source=interface.go
type DayRepoInterface interface {
	GetDays(office_id string, date string) ([]dayEntities.Day, error)
	GetDayID(dayID string) (dayEntities.Day, error)
	UpdateDay(day dayEntities.Day) (dayEntities.Day, error)
}
