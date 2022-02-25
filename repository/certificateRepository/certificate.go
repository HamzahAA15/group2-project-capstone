package certificateRepository

import (
	"database/sql"
	"log"
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

func (cr *certificateRepo) GetCertificate(id string) (certificateEntities.Certificate, error) {
	var certificate certificateEntities.Certificate

	row := cr.db.QueryRow(`SELECT id, user_id, image, dosage, status, created_at, updated_at FROM certificates WHERE id = ?`, id)

	err := row.Scan(&certificate.ID, &certificate.User.ID, &certificate.Image, &certificate.Dosage, &certificate.Status, &certificate.CreatedAt, &certificate.UpdatedAt)
	if err != nil {
		return certificate, err
	}

	return certificate, nil
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

func (cr *certificateRepo) CountVaccineIsAccept(userID string, dossage int) int {
	row, err := cr.db.Query(`
	SELECT
		COUNT(certificates.user_id)
	FROM
		certificates
	WHERE
			certificates.user_id = ?
		AND
			certificates.status = "pending"
		AND 
			certificates.dosage = ?
	ORDER BY
		certificates.created_at DESC`, userID, dossage)

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	var count int

	for row.Next() {
		if err := row.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	return count
}

func (cr *certificateRepo) GetVaccineDose(userID string) int {
	row, err := cr.db.Query(`
	SELECT
		COUNT(certificates.user_id)
	FROM
		certificates
	WHERE
			certificates.user_id = ?
		AND
			certificates.status = "approved"
	GROUP BY
		certificates.user_id`, userID)

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	var count int

	for row.Next() {
		if err := row.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	return count
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

func (cr *certificateRepo) VerifyCertificate(certificate certificateEntities.Certificate) (certificateEntities.Certificate, error) {
	query := `UPDATE certificates SET status = ?, admin_id = ?, updated_at = ? WHERE id = ?`

	statement, err := cr.db.Prepare(query)
	if err != nil {
		return certificate, err
	}

	defer statement.Close()

	_, err = statement.Exec(certificate.Status, certificate.Admin.ID, certificate.UpdatedAt, certificate.ID)
	if err != nil {
		return certificate, err
	}

	return certificate, nil
}
