package webserver

import (
	"log"
	"net/http"

	"github.com/martini"
)

var (
	listenAddress = ":8080"
)

func init() {

}

func RunServer() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello World!"
	})

	log.Printf("Listening on %s...", listenAddress)
	http.Handle("/", m)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
