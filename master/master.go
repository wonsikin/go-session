package master

import (
	"net/http"

	"github.com/omigo/log"

	"github.com/wonsikin/go-session/session"
)

var GlobalSession *session.Manager

func init() {
	GlobalSession, _ = session.NewManager("memory", "gosessionid", 3600)
	go GlobalSession.GC()
}

func Route() (interceptor *SessionInterceptor) {
	interceptor = NewSessionInterceptor()

	interceptor.HandleFunc("/master/find", findHandle)
	return interceptor
}

func findHandle(w http.ResponseWriter, r *http.Request) {
	log.Debugf("method is %s", r.Method)
	w.Write([]byte("Hello World"))
}

type SessionInterceptor struct {
	http.ServeMux
}

func NewSessionInterceptor() *SessionInterceptor {
	return &SessionInterceptor{*http.NewServeMux()}
}

func (interceptor *SessionInterceptor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := GlobalSession.SessionStart(w, r)
	log.Infof("path is %s", r.URL.Path)
	log.Infof("session id is %s", sess.SessionID())

	h, _ := interceptor.Handler(r)
	h.ServeHTTP(w, r)
}
