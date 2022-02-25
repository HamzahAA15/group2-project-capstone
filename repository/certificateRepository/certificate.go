package certificateRepository

import (
	"database/sql"
	"sirclo/project-capstone/entities/certificateEntities"
)

type certificateRepo struct {
	db *sql.DB
}

func NewMySQLCertificateRepository(db *sql.DB) CertificateInterface {
	return &certificateRepo{
		db: db,
	}
}

func (cr *certificateRepo) GetCertificates(officeID string) ([]certificateEntities.Certificate, error) {
	var certificates []certificateEntities.Certificate

	result, err := cr.db.Query(`
	SELECT 
		certificates.id, certificates.image, certificates.dosage, certificates.status, 
		user.name AS employee, (COALESCE(NULLIF(admin.name,''), '-')) AS admin
	FROM 
		certificates
	LEFT JOIN 
		users AS user ON user.id = certificates.user_id
	LEFT JOIN
		users AS admin ON admin.id = certificates.admin_id
	ORDER BY 
		certificates.updated_at DESC`)
	if err != nil {
		return certificates, err
	}

	for result.Next() {
		var certificate certificateEntities.Certificate

		errScan := result.Scan(&certificate.ID, &certificate.Image, &certificate.Dosage, &certificate.Status, &certificate.User.Name, &certificate.Admin.Name)

		if errScan != nil {
			return certificates, errScan
		}

		certificates = append(certificates, certificate)
	}

	return certificates, nil
}

func (cr *certificateRepo) GetCertificateUser(userID string) ([]certificateEntities.Certificate, error) {
	var certificates []certificateEntities.Certificate

	result, err := cr.db.Query(`
	SELECT 
		certificates.id, certificates.image, certificates.dosage, certificates.status, 
		user.name AS employee, (COALESCE(NULLIF(admin.name,''), '-')) AS admin
	FROM 
		certificates
	LEFT JOIN 
		users AS user ON user.id = certificates.user_id
	LEFT JOIN
		users AS admin ON admin.id = certificates.admin_id
	WHERE
		user.id = ?
	ORDER BY 
		certificates.updated_at DESC`, userID)
	if err != nil {
		return certificates, err
	}

	for result.Next() {
		var certificate certificateEntities.Certificate

		errScan := result.Scan(&certificate.ID, &certificate.Image, &certificate.Dosage, &certificate.Status, &certificate.User.Name, &certificate.Admin.Name)

		if errScan != nil {
			return certificates, errScan
		}

		certificates = append(certificates, certificate)
	}

	return certificates, nil
}

func (cr *certificateRepo) UploadCertificateVaccine(certificate certificateEntities.Certificate) (certificateEntities.Certificate, error) {
	query := `INSERT INTO certificates (id, user_id, image, dosage, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`

	statement, err := cr.db.Prepare(query)
	if err != nil {
		return certificate, err
	}

	defer statement.Close()

	_, err = statement.Exec(certificate.ID, certificate.User.ID, certificate.Image, certificate.Dosage, certificate.Status, certificate.CreatedAt, certificate.UpdatedAt)
	if err != nil {
		return certificate, err
	}

	return certificate, nil

}
