package certificateService

import (
	"sirclo/project-capstone/entities/certificateEntities"
	"sirclo/project-capstone/utils/request/certificateRequest"
)

type CertificateServiceInterface interface {
	GetCertificates(officeId string) ([]certificateEntities.Certificate, error)
	GetCertificate(id string) (certificateEntities.Certificate, error)
	GetCertificateUser(userId string) ([]certificateEntities.Certificate, error)
	CountVaccineAccept(userID string, dossage int) int
	GetVaccineDose(userID string) int
	UploadCertificateVaccine(userID string, input certificateRequest.CertificateUploadRequest) error
	VerifyCertificate(id string, userID string, input certificateRequest.CertificateUploadRequest) (certificateEntities.Certificate, error)
}
