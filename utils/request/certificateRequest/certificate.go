package certificateRequest

type CertificateUploadRequest struct {
	ID     string `json:"id" form:"id"`
	Image  string `json:"image" form:"image"`
	Dosage int    `json:"dosage" form:"dosage"`
	Status string `json:"status" form:"status"`
}
