package certificateService

import (
	"sirclo/project-capstone/entities/certificateEntities"
	"sirclo/project-capstone/utils/request/certificateRequest"
)

type CertificateServiceInterface interface {
	GetCertificates(officeId string) ([]certificateEntities.Certificate, error)
	GetCertificateUser(userId string) ([]certificateEntities.Certificate, error)
	UploadCertificateVaccine(userID string, input certificateRequest.CertificateUploadRequest) error
}
