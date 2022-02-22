package dayResponse

import (
	"sirclo/project-capstone/entities/dayEntities"
	"time"
)

type DayResponse struct {
	ID    string    `json:"id"`
	Date  time.Time `json:"date"`
	Quota int       `json:"Quota"`
}

type DayUpdateResponse struct {
	ID    string `json:"id"`
	Quota int    `json:"Quota"`
}

func FormatDay(day dayEntities.Day) DayResponse {
	fomatter := DayResponse{
		ID:    day.ID,
		Date:  day.Date,
		Quota: day.Quota,
	}
	return fomatter
}

func FormatUpdateDay(day dayEntities.Day) DayUpdateResponse {
	fomatter := DayUpdateResponse{
		ID:    day.ID,
		Quota: day.Quota,
	}
	return fomatter
}
