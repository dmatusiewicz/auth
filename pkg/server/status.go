package server

import (
	"net/http"

	"go.uber.org/zap"
)

func status(w http.ResponseWriter, r *http.Request) {
	requestLog = []zap.Field{
		zap.String("RemoteAddr", r.RemoteAddr),
		zap.String("Host", r.Host),
		zap.String("RequestURI", r.RequestURI),
		zap.String("Method", r.Method),
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("It works"))
	requestLog = append(requestLog, zap.Int("Status", http.StatusOK))
	logger.Debug("Status check", requestLog...)
}
