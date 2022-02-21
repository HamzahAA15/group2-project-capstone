package dayResponse

import (
	"sirclo/project-capstone/entities/dayEntities"
	"time"
)

type DayResponse struct {
	Date  time.Time
	Quota int
}

func FormatDay(day dayEntities.Day) DayResponse {
	fomatter := DayResponse{
		Date:  day.Date,
		Quota: day.Quota,
	}
	return fomatter
}
