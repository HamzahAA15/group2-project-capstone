package certificateService

import (
	"errors"
	"sirclo/project-capstone/entities/certificateEntities"
	"sirclo/project-capstone/entities/userEntities"
	"sirclo/project-capstone/utils/request/certificateRequest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetCertificates(t *testing.T) {
	certificateServiceMock := NewCertificateService(mockCertificateRepository{})
	certificate, err := certificateServiceMock.GetCertificates("ASC")
	expected, _ := []certificateEntities.Certificates{
		{
			User: userEntities.User{
				ID:    "1-user",
				Name:  "User Name",
				Email: "User Email",
			},
			Certificates: []certificateEntities.Certificate{
				{
					ID:     "1-certificate",
					Dosage: 1,
				},
			},
		},
	}, err

	assert.Nil(t, err)
	assert.Equal(t, expected, certificate)
}

func TestGetCertificate(t *testing.T) {
	certificateServiceMock := NewCertificateService(mockCertificateRepository{})
	certificate, err := certificateServiceMock.GetCertificate("1-ceritificate")
	assert.Nil(t, err)
	assert.Equal(t, "1-certificate", certificate.ID, "ID Certificate tidak sama")
	assert.Equal(t, 1, certificate.Dosage, "Dosis tidak sama")
}

func TestGetCertificateUser(t *testing.T) {
	certificateServiceMock := NewCertificateService(mockCertificateRepository{})
	cerificate, err := certificateServiceMock.GetCertificateUser("1-user")
	expected, _ := []certificateEntities.Certificate{
		{
			ID:     "1-certificate",
			Dosage: 1,
		},
		{
			ID:     "2-certificate",
			Dosage: 2,
		},
	}, err
	assert.Nil(t, err)
	assert.Equal(t, expected, cerificate)
}

func TestCountVaccineIsPending(t *testing.T) {
	certificateServiceMock := NewCertificateService(mockCertificateRepository{})
	count := certificateServiceMock.CountVaccineIsPending("1-user", 1)
	assert.Equal(t, 1, count, "tidak sesuai")
}

func TestGetVaccineDose(t *testing.T) {
	certificateServiceMock := NewCertificateService(mockCertificateRepository{})
	count := certificateServiceMock.GetVaccineDose("1-user", "approved")
	assert.Equal(t, 1, count, "tidak sesuai")
}

func TestUploadCertificateVaccine(t *testing.T) {
	certificateServiceMock := NewCertificateService(mockCertificateRepository{})

	input := certificateRequest.CertificateUploadRequest{
		ID:     "1",
		Image:  "https://urlimage",
		Dosage: 1,
		Status: "pending",
	}

	certificate, err := certificateServiceMock.UploadCertificateVaccine("1-user", input)
	assert.Nil(t, err)
	assert.Equal(t, "1-certificate", certificate.ID, "tidak sama")
	assert.Equal(t, 1, certificate.Dosage, "tidak sama")
}

func TestVerifyCertificate(t *testing.T) {
	certificateServiceMock := NewCertificateService(mockCertificateRepository{})

	input := certificateRequest.CertificateVerificationRequest{
		ID:        "1",
		Status:    "approved",
		AdminID:   "1-admin",
		UpdatedAt: time.Now(),
	}

	certificate, _ := certificateServiceMock.VerifyCertificate("1-certificate", "1-user", input)
	assert.Equal(t, "1-certificate", certificate.ID, "tidak sama")
	assert.Equal(t, "approved", certificate.Status, "tidak sama")
}

func TestVerifyCertificateError(t *testing.T) {
	certificateServiceMock := NewCertificateService(mockCertificateRepository{})
	input := certificateRequest.CertificateVerificationRequest{
		ID:        "5",
		Status:    "approved",
		AdminID:   "1-admin",
		UpdatedAt: time.Now(),
	}
	_, err := certificateServiceMock.VerifyCertificate("5", "1-user", input)
	assert.NotNil(t, err)
}

type mockCertificateRepository struct{}

func (mock mockCertificateRepository) GetCertificates(orderBy string) ([]certificateEntities.Certificates, error) {
	return []certificateEntities.Certificates{
		{
			User: userEntities.User{
				ID:    "1-user",
				Name:  "User Name",
				Email: "User Email",
			},
			Certificates: []certificateEntities.Certificate{
				{
					ID:     "1-certificate",
					Dosage: 1,
				},
			},
		},
	}, nil
}

func (mock mockCertificateRepository) GetCertificate(id string) (certificateEntities.Certificate, error) {
	if id == "5" {
		return certificateEntities.Certificate{}, errors.New("error bang")
	}
	return certificateEntities.Certificate{
		ID:     "1-certificate",
		Dosage: 1,
	}, nil
}

func (mock mockCertificateRepository) GetCertificateUser(userID string) ([]certificateEntities.Certificate, error) {
	return []certificateEntities.Certificate{
		{
			ID:     "1-certificate",
			Dosage: 1,
		},
		{
			ID:     "2-certificate",
			Dosage: 2,
		},
	}, nil
}

func (mock mockCertificateRepository) CountVaccineIsPending(userID string, dossage int) int {
	return 1
}

func (mock mockCertificateRepository) GetVaccineDose(userID string, status string) int {
	return 1
}

func (mock mockCertificateRepository) UploadCertificateVaccine(certificate certificateEntities.Certificate) (certificateEntities.Certificate, error) {
	return certificateEntities.Certificate{
		ID:     "1-certificate",
		Dosage: 1,
	}, nil
}

func (mock mockCertificateRepository) VerifyCertificate(certificate certificateEntities.Certificate) (certificateEntities.Certificate, error) {
	return certificateEntities.Certificate{
		ID:     "1-certificate",
		Dosage: 1,
		Status: "approved",
		Admin:  userEntities.User{ID: "1-admin"},
	}, nil
}
