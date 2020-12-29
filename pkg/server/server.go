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
	if port == ":" {
		log.Fatal("please provide a PORT environment variable")
	}

	fileServer := http.FileServer(http.Dir("./static"))
	// views
	r.HandleFunc("/", handlers.HtmlHandler("index"))
	r.HandleFunc("/close", handlers.HtmlHandler("close"))

	// static
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", fileServer))

	// oauth handlers
	r.HandleFunc("/credentials", handlers.StoreCredentialsHandler).Methods(http.MethodPost)
	r.HandleFunc("/token", handlers.TokenHandler)

	// api
	r.HandleFunc("/api/queue", handlers.QueueHandler).Methods(http.MethodGet)

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(port, r))
}
