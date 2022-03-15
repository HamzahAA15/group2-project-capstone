package officeRepository

import (
	"sirclo/project-capstone/entities/officeEntities"
)

//go:generate mockgen --destination=./../../mocks/office/repository/mock_repository_office.go -source=interface.go
type OfficeRepoInterface interface {
	GetOffices() ([]officeEntities.Office, error)
}
