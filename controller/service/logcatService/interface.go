package logcatService

import "sirclo/project-capstone/entities/logcatEntities"

type LogcatServiceInterface interface {
	CreateLogcat(loginId, message, category string) (logcatEntities.Logcat, error)
	GetLogcats() ([]logcatEntities.Logcat, error)
	GetLogcatUser(userID string) ([]logcatEntities.Logcat, error)
}
