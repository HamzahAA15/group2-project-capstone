package dayResponse

import (
	"sirclo/project-capstone/entities/dayEntities"
	"time"
)

type DayResponse struct {
	ID    string
	Date  time.Time
	Quota int
}

func FormatDay(day dayEntities.Day) DayResponse {
	fomatter := DayResponse{
		ID:    day.ID,
		Date:  day.Date,
		Quota: day.Quota,
	}
	return fomatter
}
