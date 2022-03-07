package logcatHandler

import "net/http"

type LogcatHandlerInterface interface {
	LogcatsHandler(w http.ResponseWriter, r *http.Request)
	LogcatUserHandler(w http.ResponseWriter, r *http.Request)
}
