package logcatResponse

import (
	"sirclo/project-capstone/entities/logcatEntities"
	"time"
)

type LogcatResponse struct {
	ID        string    `json:"id"`
	Message   string    `json:"name"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatLog(log logcatEntities.Logcat) LogcatResponse {
	formatter := LogcatResponse{
		ID:        log.ID,
		Message:   log.Message,
		Category:  log.Category,
		CreatedAt: log.CreatedAt,
	}

	return formatter
}
