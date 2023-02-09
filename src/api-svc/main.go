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
	service.SetupRoutes(r, svc)

	log.Println("Started server at", *httpAddr)
	err := http.ListenAndServe(*httpAddr, r)
	if err != nil {
		log.Fatalln("error while starting the server. ", err)
	}
}
