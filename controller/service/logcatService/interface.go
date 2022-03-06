package logcatService

import "sirclo/project-capstone/entities/logcatEntities"

type LogcatServiceInterface interface {
	CreateLogcat(loginId, message, category string) (logcatEntities.Logcat, error)
}
