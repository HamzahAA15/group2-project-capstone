package certificateResponse

import (
	"sirclo/project-capstone/entities/certificateEntities"
)

type CertificateResponse struct {
	ID     string `json:"id"`
	User   string `json:"user"`
	Image  string `json:"image"`
	Dosage int    `json:"dosage"`
	Status string `json:"status"`
	Admin  string `json:"admin"`
}

func FormatCertificate(certificate certificateEntities.Certificate) CertificateResponse {
	fomatter := CertificateResponse{
		ID:     certificate.ID,
		User:   certificate.User.Name,
		Image:  certificate.Image,
		Dosage: certificate.Dosage,
		Status: certificate.Status,
		Admin:  certificate.Admin.Name,
	}

	return fomatter
}

type CertificatesResponse struct {
	UserID       string                `json:"user_id"`
	UserName     string                `json:"user_name"`
	Certificates []CertificateResponse `json:"certificates"`
}

func FormatCertificates(certificate certificateEntities.Certificates) CertificatesResponse {
	data := []CertificateResponse{}

	for i := range certificate.Certificates {
		certificateFormatter := FormatCertificate(certificate.Certificates[i])
		data = append(data, certificateFormatter)
	}

	formatter := CertificatesResponse{
		UserID:       certificate.User.ID,
		UserName:     certificate.User.Name,
		Certificates: data,
	}

	return formatter
}
