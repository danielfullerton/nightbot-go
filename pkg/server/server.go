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
	r.HandleFunc("/", handlers.HtmlHandler("index"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", fileServer))

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(port, r))
}
