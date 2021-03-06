package logcatService

import (
	"sirclo/project-capstone/entities/logcatEntities"
	"sirclo/project-capstone/repository/logcatRepository"

	"github.com/google/uuid"
)

type logcatService struct {
	lcRepo logcatRepository.LogcatRepoInterface
}

func NewLogcatService(lcRepo logcatRepository.LogcatRepoInterface) LogcatServiceInterface {
	return &logcatService{
		lcRepo: lcRepo,
	}
}

func (ls *logcatService) CreateLogcat(loginId, message, category string) (logcatEntities.Logcat, error) {
	var logcat logcatEntities.Logcat
	logcat.ID = uuid.New().String()
	logcat.User.ID = loginId
	logcat.Message = message
	logcat.Category = category

	createLogcat, err := ls.lcRepo.CreateLogcat(logcat)
	return createLogcat, err
}

func (ls *logcatService) GetLogcats() ([]logcatEntities.Logcat, error) {
	logcats, err := ls.lcRepo.GetLogcats()
	return logcats, err
}

func (ls *logcatService) GetLogcatUser(userID string) ([]logcatEntities.Logcat, error) {
	logcats, err := ls.lcRepo.GetLogcatUser(userID)
	return logcats, err
}
