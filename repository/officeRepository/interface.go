package officeRepository

import (
	"sirclo/project-capstone/entities/officeEntities"
)

type OfficeRepoInterface interface {
	GetOffices() ([]officeEntities.Office, error)
}
