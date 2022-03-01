package dayResponse

import (
	"sirclo/project-capstone/entities/dayEntities"
	"time"
)

type DayResponse struct {
	ID             string    `json:"id"`
	Office         string    `json:"office"`
	Date           time.Time `json:"date"`
	Quota          int       `json:"quota"`
	TotalApproved  int       `json:"total_approved"`
	RemainingQuota int       `json:"remaining_quota"`
}

type DayUpdateResponse struct {
	ID    string `json:"id"`
	Quota int    `json:"quota"`
}

func FormatDay(day dayEntities.Day) DayResponse {
	fomatter := DayResponse{
		ID:             day.ID,
		Office:         day.OfficeId.Name,
		Date:           day.Date,
		Quota:          day.Quota,
		TotalApproved:  day.TotalApproved,
		RemainingQuota: day.RemainingQuota,
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
