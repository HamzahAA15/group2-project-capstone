package logcatRepository

import "sirclo/project-capstone/entities/logcatEntities"

//go:generate mockgen --destination=./../../mocks/logcat/repository/mock_repository_logcat.go -source=interface.go
type LogcatRepoInterface interface {
	CreateLogcat(lc logcatEntities.Logcat) (logcatEntities.Logcat, error)
	GetLogcats() ([]logcatEntities.Logcat, error)
	GetLogcatUser(userID string) ([]logcatEntities.Logcat, error)
}
