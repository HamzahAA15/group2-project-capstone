package dayResponse

import (
	"sirclo/project-capstone/entities/dayEntities"
	"time"
)

type DayResponse struct {
<<<<<<< HEAD
	ID    string
	Date  time.Time
	Quota int
=======
	ID    string    `json:"id"`
	Date  time.Time `json:"date"`
	Quota int       `json:"Quota"`
}

type DayUpdateResponse struct {
	ID    string `json:"id"`
	Quota int    `json:"Quota"`
>>>>>>> 26a06d74ef39e7335a5b239eab9352123ca496bb
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
