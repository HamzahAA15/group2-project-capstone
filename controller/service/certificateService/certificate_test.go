package certificateService

import (
	"errors"
	"testing"
	"time"

	_certificateEntities "sirclo/project-capstone/entities/certificateEntities"
	_userEntities "sirclo/project-capstone/entities/userEntities"
	_mockCertificateRepo "sirclo/project-capstone/mocks/certificate/repository"
	_utils "sirclo/project-capstone/utils"
	_certificateRequest "sirclo/project-capstone/utils/request/certificateRequest"

	gomock "github.com/golang/mock/gomock"
)

func TestGetCertificates(t *testing.T) {
	certificatesData := dummyCertificates(t)

	testCases := []struct {
		name                string
		orderBy             string
		mockCertificateRepo func(certificate *_mockCertificateRepo.MockCertificateInterface)
	}{
		{
			name:    "success",
			orderBy: "asc",
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().GetCertificates(gomock.Eq("asc")).Return(certificatesData, nil).Times(1)
			},
		},
		{
			name:    "error",
			orderBy: "desc",
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().GetCertificates(gomock.Eq("desc")).Return(certificatesData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockCertificateRepo.NewMockCertificateInterface(ctrl)
			tc.mockCertificateRepo(mockRepo)

			s := NewCertificateService(mockRepo)
			_, _ = s.GetCertificates(tc.orderBy)
		})
	}
}

func TestGetCertificate(t *testing.T) {
	certificateData := dummyCertificate(t)

	testCases := []struct {
		name                string
		id                  string
		mockCertificateRepo func(certificate *_mockCertificateRepo.MockCertificateInterface)
	}{
		{
			name: "success",
			id:   "1",
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().GetCertificate(gomock.Eq("1")).Return(certificateData, nil).Times(1)
			},
		},
		{
			name: "error",
			id:   "2",
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().GetCertificate(gomock.Eq("2")).Return(certificateData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockCertificateRepo.NewMockCertificateInterface(ctrl)
			tc.mockCertificateRepo(mockRepo)

			s := NewCertificateService(mockRepo)
			_, _ = s.GetCertificate(tc.id)
		})
	}
}

func TestGetCertificateUser(t *testing.T) {
	certificatesData := dummyCertificateUser(t)

	testCases := []struct {
		name                string
		idUser              string
		mockCertificateRepo func(certificate *_mockCertificateRepo.MockCertificateInterface)
	}{
		{
			name:   "success",
			idUser: "1",
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().GetCertificateUser(gomock.Eq("1")).Return(certificatesData, nil).Times(1)
			},
		},
		{
			name:   "error",
			idUser: "2",
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().GetCertificateUser(gomock.Eq("2")).Return(certificatesData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockCertificateRepo.NewMockCertificateInterface(ctrl)
			tc.mockCertificateRepo(mockRepo)

			s := NewCertificateService(mockRepo)
			_, _ = s.GetCertificateUser(tc.idUser)
		})
	}
}

func TestCountVaccineIsPending(t *testing.T) {
	testCases := []struct {
		name                string
		idUser              string
		dose                int
		mockCertificateRepo func(certificate *_mockCertificateRepo.MockCertificateInterface)
	}{
		{
			name:   "success",
			idUser: "1",
			dose:   1,
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().CountVaccineIsPending(gomock.Eq("1"), gomock.Eq(1)).Return(1).Times(1)
			},
		},
		{
			name:   "error",
			idUser: "2",
			dose:   2,
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().CountVaccineIsPending(gomock.Eq("2"), gomock.Eq(2)).Return(0).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockCertificateRepo.NewMockCertificateInterface(ctrl)
			tc.mockCertificateRepo(mockRepo)

			s := NewCertificateService(mockRepo)
			_ = s.CountVaccineIsPending(tc.idUser, tc.dose)
		})
	}
}

func TestGetVaccineDose(t *testing.T) {
	testCases := []struct {
		name                string
		idUser              string
		status              string
		mockCertificateRepo func(certificate *_mockCertificateRepo.MockCertificateInterface)
	}{
		{
			name:   "success",
			idUser: "1",
			status: "approved",
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().GetVaccineDose(gomock.Eq("1"), gomock.Eq("approved")).Return(1).Times(1)
			},
		},
		{
			name:   "error",
			idUser: "2",
			status: "approved",
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().GetVaccineDose(gomock.Eq("2"), gomock.Eq("approved")).Return(0).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockCertificateRepo.NewMockCertificateInterface(ctrl)
			tc.mockCertificateRepo(mockRepo)

			s := NewCertificateService(mockRepo)
			_ = s.GetVaccineDose(tc.idUser, tc.status)
		})
	}
}

func TestUploadCertificateVaccine(t *testing.T) {
	certificateData := dummyCertificate(t)

	testCases := []struct {
		name                string
		idUser              string
		body                _certificateRequest.CertificateUploadRequest
		mockCertificateRepo func(certificate *_mockCertificateRepo.MockCertificateInterface)
	}{
		{
			name:   "success",
			idUser: "1",
			body: _certificateRequest.CertificateUploadRequest{
				Image: _utils.RandomString(6),
			},
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().UploadCertificateVaccine(gomock.Any()).Return(certificateData, nil).Times(1)
			},
		},
		{
			name:   "error",
			idUser: "2",
			body: _certificateRequest.CertificateUploadRequest{
				Image: _utils.RandomString(6),
			},
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().UploadCertificateVaccine(gomock.Any()).Return(certificateData, errors.New("error")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockCertificateRepo.NewMockCertificateInterface(ctrl)
			tc.mockCertificateRepo(mockRepo)

			s := NewCertificateService(mockRepo)
			_, _ = s.UploadCertificateVaccine(tc.idUser, tc.body)
		})
	}
}

func TestVerifyCertificate(t *testing.T) {
	certificateData := dummyCertificate(t)

	testCases := []struct {
		name                string
		id                  string
		idUser              string
		body                _certificateRequest.CertificateVerificationRequest
		mockCertificateRepo func(certificate *_mockCertificateRepo.MockCertificateInterface)
	}{
		{
			name:   "success",
			id:     "1",
			idUser: "1",
			body: _certificateRequest.CertificateVerificationRequest{
				ID:        "1",
				Status:    "approved",
				AdminID:   _utils.RandomString(8),
				UpdatedAt: time.Now(),
			},
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().GetCertificate(gomock.Eq("1")).Return(certificateData, nil).Times(1)
				certificate.EXPECT().VerifyCertificate(gomock.Any()).Return(certificateData, nil).Times(1)
			},
		},
		{
			name:   "error ID",
			id:     "2",
			idUser: "2",
			body: _certificateRequest.CertificateVerificationRequest{
				ID:        "2",
				Status:    "approved",
				AdminID:   _utils.RandomString(8),
				UpdatedAt: time.Now(),
			},
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().GetCertificate(gomock.Eq("2")).Return(certificateData, errors.New("error ID")).Times(1)
			},
		},
		{
			name:   "error verification",
			id:     "3",
			idUser: "3",
			body: _certificateRequest.CertificateVerificationRequest{
				ID:        "3",
				Status:    "approved",
				AdminID:   _utils.RandomString(8),
				UpdatedAt: time.Now(),
			},
			mockCertificateRepo: func(certificate *_mockCertificateRepo.MockCertificateInterface) {
				certificate.EXPECT().GetCertificate(gomock.Eq("3")).Return(certificateData, nil).Times(1)
				certificate.EXPECT().VerifyCertificate(gomock.Any()).Return(certificateData, errors.New("error verification")).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := _mockCertificateRepo.NewMockCertificateInterface(ctrl)
			tc.mockCertificateRepo(mockRepo)

			s := NewCertificateService(mockRepo)
			_, _ = s.VerifyCertificate(tc.id, tc.idUser, tc.body)
		})
	}
}

func dummyCertificates(t *testing.T) (certificates []_certificateEntities.Certificates) {
	certificates = []_certificateEntities.Certificates{
		{
			User: _userEntities.User{
				ID:    _utils.RandomString(8),
				Name:  _utils.RandomString(10),
				Email: _utils.RandomString(4) + "@" + _utils.RandomString(5) + ".com",
			},
			Certificates: []_certificateEntities.Certificate{
				{
					ID:     _utils.RandomString(8),
					Image:  "https://" + _utils.RandomString(4) + ".com",
					Dosage: 1,
					Status: "pending",
					Admin:  _userEntities.User{ID: _utils.RandomString(8)},
				},
			},
		},
	}

	return
}

func dummyCertificateUser(t *testing.T) (certificates []_certificateEntities.Certificate) {
	certificates = []_certificateEntities.Certificate{
		{
			ID:     _utils.RandomString(8),
			Image:  "https://" + _utils.RandomString(4) + ".com",
			Dosage: 1,
			Status: "approved",
			Admin:  _userEntities.User{ID: _utils.RandomString(8)},
		},
	}

	return
}

func dummyCertificate(t *testing.T) (certificate _certificateEntities.Certificate) {
	certificate = _certificateEntities.Certificate{
		ID:     _utils.RandomString(8),
		Image:  "https://" + _utils.RandomString(4) + ".com",
		Dosage: 1,
		Status: "pending",
		Admin:  _userEntities.User{ID: _utils.RandomString(8)},
	}

	return
}
