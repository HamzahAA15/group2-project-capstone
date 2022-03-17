package certificateRepository

import "sirclo/project-capstone/entities/certificateEntities"

//go:generate mockgen --destination=./../../mocks/certificate/repository/mock_repository_certificate.go -source=interface.go
type CertificateInterface interface {
	GetCertificates(orderBy string) ([]certificateEntities.Certificates, error)
	GetCertificate(id string) (certificateEntities.Certificate, error)
	GetCertificateUser(userID string) ([]certificateEntities.Certificate, error)
	CountVaccineIsPending(userID string, dossage int) int
	GetVaccineDose(userID string, status string) int
	UploadCertificateVaccine(certificate certificateEntities.Certificate) (certificateEntities.Certificate, error)
	VerifyCertificate(certificate certificateEntities.Certificate) (certificateEntities.Certificate, error)
}
