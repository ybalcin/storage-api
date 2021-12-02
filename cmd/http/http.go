package http

import (
	"fmt"
	"github.com/ybalcin/storage-api/internal/port"
	"log"
	"net/http"
	"os"
)

// Serve starts the http server
func Serve() {
	httpPort := port.NewHttpServer()

	mux := http.NewServeMux()

	mux.Handle("/", port.Handler{H: httpPort.GetRecordsHandler, Methods: map[string]struct{}{http.MethodPost: {}}})
	mux.Handle("/in-memory", port.Handler{H: httpPort.CacheEntryGroupHandler, Methods: map[string]struct{}{http.MethodGet: {}, http.MethodPost: {}}})

	p, isExist := os.LookupEnv("PORT")
	if !isExist {
		p = "8080"
	}

	log.Printf("http server listening on port: %s", p)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", p), mux))
}
