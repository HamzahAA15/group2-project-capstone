package certificateService

import (
	"fmt"
	"sirclo/project-capstone/entities/certificateEntities"
	"sirclo/project-capstone/repository/certificateRepository"
	"sirclo/project-capstone/utils/request/certificateRequest"
	"time"

	"github.com/google/uuid"
)

type certificateService struct {
	certificateRepository certificateRepository.CertificateInterface
}

func NewCertificateService(repo certificateRepository.CertificateInterface) CertificateServiceInterface {
	return &certificateService{
		certificateRepository: repo,
	}
}

func (cs *certificateService) GetCertificates(officeID string) ([]certificateEntities.Certificate, error) {
	certificates, err := cs.certificateRepository.GetCertificates(officeID)
	return certificates, err
}

func (cs *certificateService) GetCertificate(id string) (certificateEntities.Certificate, error) {
	certificate, err := cs.certificateRepository.GetCertificate(id)
	return certificate, err
}

func (cs *certificateService) GetCertificateUser(userID string) ([]certificateEntities.Certificate, error) {
	certificates, err := cs.certificateRepository.GetCertificateUser(userID)
	return certificates, err
}

func (cs *certificateService) CountVaccineAccept(userID string, dossage int) int {
	countData := cs.certificateRepository.CountVaccineIsAccept(userID, dossage)
	return countData
}

func (cs *certificateService) GetVaccineDose(userID string) int {
	countData := cs.certificateRepository.GetVaccineDose(userID)
	return countData
}

func (cs *certificateService) UploadCertificateVaccine(userID string, input certificateRequest.CertificateUploadRequest) error {
	upload := certificateEntities.Certificate{}
	upload.ID = uuid.New().String()
	upload.User.ID = userID
	upload.Image = input.Image
	upload.Dosage = input.Dosage
	upload.Status = "pending"
	upload.CreatedAt = time.Now()
	upload.UpdatedAt = time.Now()

	_, err := cs.certificateRepository.UploadCertificateVaccine(upload)
	return err
}

func (cs *certificateService) VerifyCertificate(id string, userID string, input certificateRequest.CertificateUploadRequest) (certificateEntities.Certificate, error) {
	certificate, _ := cs.GetCertificate(id)

	certificate.ID = id
	certificate.Status = input.Status
	certificate.Admin.ID = userID
	certificate.UpdatedAt = time.Now()

	fmt.Println("ini certificate", certificate)

	updateCertificate, err := cs.certificateRepository.VerifyCertificate(certificate)
	fmt.Println("ini err", err)
	return updateCertificate, err
}
