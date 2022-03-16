package logcatService

import "sirclo/project-capstone/entities/logcatEntities"

//go:generate mockgen --destination=./../../../mocks/logcat/service/mock_service_logcat.go -source=interface.go
type LogcatServiceInterface interface {
	CreateLogcat(loginId, message, category string) (logcatEntities.Logcat, error)
	GetLogcats() ([]logcatEntities.Logcat, error)
	GetLogcatUser(userID string) ([]logcatEntities.Logcat, error)
}
