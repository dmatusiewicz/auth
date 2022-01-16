package server

import (
	"net/http"

	"go.uber.org/zap"
)

func logout(w http.ResponseWriter, r *http.Request) {
	requestLog = []zap.Field{
		zap.String("RemoteAddr", r.RemoteAddr),
		zap.String("Host", r.Host),
		zap.String("RequestURI", r.RequestURI),
		zap.String("Method", r.Method),
	}
	session, _ := sessionStore.Get(r, "auth")
	if session.Values["authenticated"] == true {
		session.Values["authenticated"] = false
		session.Save(r, w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logged out"))
		username := session.Values["username"]
		requestLog = append(requestLog, zap.Int("Status", http.StatusOK))
		logger.Debug("Closing session for user: "+username.(string), requestLog...)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		requestLog = append(requestLog, zap.Int("Status", http.StatusBadRequest))
		logger.Warn("Session is alread inacvite", requestLog...)
	}
}
