package http

import (
	"fmt"
	"github.com/ybalcin/storage-api/internal/port"
	"log"
	"net/http"
)

const (
	p string = "8080"
)

// Serve starts the http server
func Serve() {
	httpPort := port.NewHttpServer()

	mux := http.NewServeMux()

	mux.Handle("/db", port.Handler{H: httpPort.GetRecordsHandler, Method: http.MethodPost})
	mux.Handle("in-memory/", port.Handler{H: httpPort.GetCacheEntryHandler, Method: http.MethodGet})
	mux.Handle("/in-memory", port.Handler{H: httpPort.AddCacheEntryHandler, Method: http.MethodPost})

	log.Printf("http server listening on port: %s", p)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", p), mux))
}
