package officeService

import (
	"sirclo/project-capstone/entities/officeEntities"
	"sirclo/project-capstone/repository/officeRepository"
)

type officeService struct {
	officeRepository officeRepository.OfficeRepoInterface
}

func NewOfficeService(repo officeRepository.OfficeRepoInterface) OfficeServiceInterface {
	return &officeService{
		officeRepository: repo,
	}
}

func (os *officeService) GetOffices() ([]officeEntities.Office, error) {
	offices, err := os.officeRepository.GetOffices()
	return offices, err
}
