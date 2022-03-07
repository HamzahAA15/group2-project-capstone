package logcatRepository

import "sirclo/project-capstone/entities/logcatEntities"

type LogcatRepoInterface interface {
	CreateLogcat(lc logcatEntities.Logcat) (logcatEntities.Logcat, error)
	GetLogcats() ([]logcatEntities.Logcat, error)
	GetLogcatUser(userID string) ([]logcatEntities.Logcat, error)
}
