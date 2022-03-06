package logcatRepository

import "sirclo/project-capstone/entities/logcatEntities"

type LogcatRepoInterface interface {
	CreateLogcat(lc logcatEntities.Logcat) (logcatEntities.Logcat, error)
}
