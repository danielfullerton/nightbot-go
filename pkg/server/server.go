package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"nightbot-go/pkg/server/handlers"
	"os"
)

var r = mux.NewRouter()

func Start() {
	port := ":" + os.Getenv("PORT")
	staticFileServer := http.FileServer(http.Dir("./static"))
	appFileServer := http.FileServer(http.Dir("./dist"))

	// views
	r.HandleFunc("/", handlers.HtmlHandler)

	// static
	r.PathPrefix("/").Handler(http.StripPrefix("/", appFileServer))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", staticFileServer))

	// oauth handlers
	r.HandleFunc("/credentials", handlers.StoreCredentialsHandler).Methods(http.MethodPost)
	r.HandleFunc("/token", handlers.TokenHandler)

	// api
	r.HandleFunc("/api/queue", handlers.QueueHandler).Methods(http.MethodGet)

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(port, r))
}
