package certificateService

import (
	"sirclo/project-capstone/entities/certificateEntities"
	"sirclo/project-capstone/repository/certificateRepository"
)

type certificateService struct {
	certificateRepository certificateRepository.CertificateInterface
}

func NewCertificateService(repo certificateRepository.CertificateInterface) CertificateServiceInterface {
	return &certificateService{
		certificateRepository: repo,
	}
}

func (cs *certificateService) GetCertificates() ([]certificateEntities.Certificate, error) {
	certificates, err := cs.certificateRepository.GetCertificates()
	return certificates, err
}
