package officeService

import "sirclo/project-capstone/entities/officeEntities"

type OfficeServiceInterface interface {
	GetOffices() ([]officeEntities.Office, error)
}
