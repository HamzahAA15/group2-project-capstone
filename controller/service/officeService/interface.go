package officeService

import "sirclo/project-capstone/entities/officeEntities"

//go:generate mockgen --destination=./../../../mocks/office/service/mock_service_office.go -source=interface.go
type OfficeServiceInterface interface {
	GetOffices() ([]officeEntities.Office, error)
}
