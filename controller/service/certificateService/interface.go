package certificateService

import "sirclo/project-capstone/entities/certificateEntities"

type CertificateServiceInterface interface {
	GetCertificates() ([]certificateEntities.Certificate, error)
}
