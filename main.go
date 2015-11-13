package main

import (
	"net/http"

	"github.com/omigo/log"

	"github.com/wonsikin/go-session/master"
)

func main() {
	port := ":9002"
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/master/", master.Route())

	log.Infof("Server is listening, addr=%s", port)
	log.Error(http.ListenAndServe(port, nil))
}
