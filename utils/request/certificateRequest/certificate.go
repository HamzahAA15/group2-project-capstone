package certificateRequest

import "time"

type CertificateUploadRequest struct {
	ID     string `json:"id" form:"id"`
	Image  string `json:"image" form:"image"`
	Dosage int    `json:"dosage" form:"dosage"`
	Status string `json:"status" form:"status"`
}

type CertificateVerificationRequest struct {
	ID        string    `json:"id" form:"id"`
	Status    string    `json:"status" form:"status"`
	AdminID   string    `json:"admin_id" form:"admin_id"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}
