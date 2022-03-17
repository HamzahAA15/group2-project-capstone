package certificateService

import (
	"sirclo/project-capstone/entities/certificateEntities"
	"sirclo/project-capstone/utils/request/certificateRequest"
)

//go:generate mockgen --destination=./../../../mocks/certificate/service/mock_service_certificate.go -source=interface.go
type CertificateServiceInterface interface {
	GetCertificates(orderBy string) ([]certificateEntities.Certificates, error)
	GetCertificate(id string) (certificateEntities.Certificate, error)
	GetCertificateUser(userId string) ([]certificateEntities.Certificate, error)
	CountVaccineIsPending(userID string, dossage int) int
	GetVaccineDose(userID string, status string) int
	UploadCertificateVaccine(userID string, input certificateRequest.CertificateUploadRequest) (certificateEntities.Certificate, error)
	VerifyCertificate(id string, userID string, input certificateRequest.CertificateVerificationRequest) (certificateEntities.Certificate, error)
}
