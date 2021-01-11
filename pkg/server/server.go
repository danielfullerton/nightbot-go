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
	staticFileServer := http.FileServer(http.Dir("./dist/assets"))
	distFileServer := http.FileServer(http.Dir("./dist"))

	// oauth handlers
	r.HandleFunc("/api/config/credentials", handlers.StoreCredentialsHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/config/token", handlers.TokenHandler)
	r.HandleFunc("/authorize", handlers.ManualAuthHandler)

	// api
	r.HandleFunc("/api/queue", handlers.QueueHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/channel", handlers.ChannelHandler).Methods(http.MethodGet)

	// static
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", staticFileServer))
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", distFileServer))

	// Serve index page on all unhandled routes
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(port, r))
}
