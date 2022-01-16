package server

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
)

// TODO: secure string should be passed from configuration or envrionment.
var codecs = securecookie.CodecsFromPairs([]byte("thisStringWillBeConfigParameter"))
var requestLog []zap.Field

// TODO: implement real session store.
// var sessionStore = sessions.NewCookieStore([]byte("test"))
var sessionStore = sessions.CookieStore{
	Options: &sessions.Options{
		HttpOnly: true,
		Path:     "/",
		MaxAge:   3600,
	},
	Codecs: codecs,
}

// TODO: implement real user store.
var userStore = map[string]string{
	"dawidm": "password1",
	"ewa":    "password1",
}

func login(w http.ResponseWriter, r *http.Request) {
	requestLog = []zap.Field{
		zap.String("Src", r.RemoteAddr),
		zap.String("Dest", r.Host),
		zap.String("RequestURI", r.RequestURI),
		zap.String("Method", r.Method),
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	session, _ := sessionStore.Get(r, "auth")
	session.Values["username"] = username
	if pass, ok := userStore[username]; ok {
		if pass == password {
			loginResponse(true, session, w, r, "Authentication is successful. User and password match for user: "+username)
		} else {
			loginResponse(false, session, w, r, "Authentication is un-successful. Wrong password.")
		}
	} else {
		loginResponse(false, session, w, r, "Authentication is un-successful. Wrong user: "+username)
	}
}

func loginResponse(auth bool, session *sessions.Session, w http.ResponseWriter, r *http.Request, message string) {
	session.Values["authenticated"] = auth
	session.Save(r, w)
	if auth {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))
		requestLog = append(requestLog, zap.Int("Status", http.StatusOK))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
		requestLog = append(requestLog, zap.Int("Status", http.StatusUnauthorized))
	}
	logger.Debug(message, requestLog...)
}
