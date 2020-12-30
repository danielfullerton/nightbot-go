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
	distFileServer := http.FileServer(http.Dir("./dist"))

	// views
	//r.HandleFunc("/", handlers.HtmlHandler)

	// oauth handlers
	r.HandleFunc("/credentials", handlers.StoreCredentialsHandler).Methods(http.MethodPost)
	r.HandleFunc("/token", handlers.TokenHandler)

	// api
	r.HandleFunc("/api/queue", handlers.QueueHandler).Methods(http.MethodGet)

	// static
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", staticFileServer))
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", distFileServer))

	// Serve index page on all unhandled routes
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(port, r))
}
