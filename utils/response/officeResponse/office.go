package officeResponse

import (
	"sirclo/project-capstone/entities/officeEntities"
)

type OfficeResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func FormatOffice(office officeEntities.Office) OfficeResponse {
	formatter := OfficeResponse{
		ID:   office.ID,
		Name: office.Name,
	}

	return formatter
}
