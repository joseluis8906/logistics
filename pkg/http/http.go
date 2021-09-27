package http

import (
	"encoding/json"
	stdhttp "net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type (
	ErrorResponse struct {
		Timestamp time.Time `json:"timestamp"`
		Status    int       `json:"status"`
		Error     string    `json:"error"`
		Message   string    `json:"message"`
		Path      string    `json:"path"`
	}
)

func Error(w stdhttp.ResponseWriter, r *stdhttp.Request, status int, resp error) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(ErrorResponse{
		Timestamp: time.Now(),
		Status:    status,
		Error:     stdhttp.StatusText(status),
		Message:   resp.Error(),
		Path:      r.URL.Path,
	})
	if err != nil {
		log.Error(err)
	}
}

func Success(w stdhttp.ResponseWriter, r *stdhttp.Request, status int, body interface{}) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Error(err)
	}
}

func Health(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	log.Debug("debug")
	log.Info("info")
	log.Warn("warning")
	log.Error("error")
	_, err := w.Write([]byte("up and running\n"))
	if err != nil {
		log.Error(err)
	}
}
