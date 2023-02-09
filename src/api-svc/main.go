package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/query-yt/src/api-svc/internal/service"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8000", "Address for HTTP server")
	)
	flag.Parse()

	r := mux.NewRouter()
	svc := service.NewHTTPServer()

	r.HandleFunc("/api/v1/videoinfo", svc.GetVideoInfo).Methods("GET")

	log.Println("Started server at", *httpAddr)
	err := http.ListenAndServe(*httpAddr, r)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
