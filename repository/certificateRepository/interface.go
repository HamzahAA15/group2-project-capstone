package certificateRepository

import "sirclo/project-capstone/entities/certificateEntities"

type CertificateInterface interface {
	GetCertificates(officeID string) ([]certificateEntities.Certificate, error)
	GetCertificateUser(userID string) ([]certificateEntities.Certificate, error)
	CountVaccineIsAccept(userID string, dossage int) int
	GetVaccineDose(userID string) int
	UploadCertificateVaccine(certificate certificateEntities.Certificate) (certificateEntities.Certificate, error)
}
