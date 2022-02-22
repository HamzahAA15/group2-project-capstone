package certificateService

import "sirclo/project-capstone/entities/certificateEntities"

type CertificateServiceInterface interface {
	GetCertificates(officeId string) ([]certificateEntities.Certificate, error)
}
