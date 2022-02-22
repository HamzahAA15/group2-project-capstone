package certificateRepository

import "sirclo/project-capstone/entities/certificateEntities"

type CertificateInterface interface {
	GetCertificates(officeID string) ([]certificateEntities.Certificate, error)
}
