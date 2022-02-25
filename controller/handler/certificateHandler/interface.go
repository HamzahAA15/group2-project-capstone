package certificateHandler

import "net/http"

type CertificateHandlerInterface interface {
	GetCertificatesHandler(w http.ResponseWriter, r *http.Request)
	GetCertificateUserHandler(w http.ResponseWriter, r *http.Request)
	UploadCertificateHandler(w http.ResponseWriter, r *http.Request)
	VerifyCertificateHandler(w http.ResponseWriter, r *http.Request)
}
