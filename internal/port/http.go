package port

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ybalcin/storage-api/internal/application"
	"github.com/ybalcin/storage-api/internal/application/commandhandler"
	"github.com/ybalcin/storage-api/internal/application/queryhandler"
	"github.com/ybalcin/storage-api/internal/common"
	"log"
	"net/http"
	"os"
)

type (
	httpServer struct {
		Application *application.Application
	}

	Handler struct {
		H       func(rw http.ResponseWriter, req *http.Request) error
		Methods map[string]struct{}
	}
)

const (
	contentType     = "Content-Type"
	applicationJson = "application/json"
	requestId       = "X-Request-ID"
)

var (
	logger = log.New(os.Stdout, "in: ", log.LstdFlags)
)

// NewHttpServer inits http server
func NewHttpServer() *httpServer {
	app := application.New()

	return &httpServer{app}
}

// NewHttpServerWithApplication initializes new httpserver with application argument
func NewHttpServerWithApplication(application *application.Application) *httpServer {
	return &httpServer{Application: application}
}

// ServeHTTP middleware
func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// log incoming request as stdout
	defer logRequest(req)

	_, ok := h.Methods[req.Method]
	if !ok {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set(contentType, applicationJson)

	err := h.H(w, req)

	if err == nil {
		if e, ok := recover().(error); ok {
			err = e
		}
	}

	if err != nil {
		switch e := err.(type) {
		case common.StatusError:
			logger.Printf("HTTP status: %d - Message: %s", e.Status(), e.Error())
			http.Error(w, e.Error(), e.Status())
		default:
			fmt.Println(e)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

// AddCacheEntryHandler adds cache entry to in memory
func (s *httpServer) AddCacheEntryHandler(w http.ResponseWriter, req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	var command commandhandler.AddCacheEntryCommand

	if err := decoder.Decode(&command); err != nil {
		return common.ThrowBadRequestError(err)
	}

	err := s.Application.Commands().AddCacheEntryCommand.Handle(&command)
	if err != nil {
		return common.ThrowBadRequestError(errors.New(err.String()))
	}

	return nil
}

// CacheEntryGroupHandler handles cache entry group
func (s *httpServer) CacheEntryGroupHandler(w http.ResponseWriter, req *http.Request) error {
	if req.Method == http.MethodGet {
		return s.GetCacheEntryHandler(w, req)
	}
	if req.Method == http.MethodPost {
		return s.AddCacheEntryHandler(w, req)
	}

	return nil
}

// GetCacheEntryHandler gets entry from cache
func (s *httpServer) GetCacheEntryHandler(w http.ResponseWriter, req *http.Request) error {
	key := req.URL.Query().Get("key")

	queryResult, err := s.Application.Queries().GetCacheEntryQuery.Handle(&queryhandler.GetCacheEntryQuery{Key: key})
	if err != nil {
		return common.ThrowBadRequestError(errors.New(err.String()))
	}

	if queryResult == nil {
		http.Error(w, "", http.StatusNoContent)
		return nil
	}

	if err := json.NewEncoder(w).Encode(queryResult); err != nil {
		return common.ThrowBadRequestError(err)
	}
	return nil
}

// GetRecordsHandler gets records
func (s *httpServer) GetRecordsHandler(w http.ResponseWriter, req *http.Request) error {
	decoder := json.NewDecoder(req.Body)

	var query queryhandler.GetRecordsQuery

	recordResponse := RecordHttpResponse{
		Code:    0,
		Message: "Success",
		Records: nil,
	}

	err := decoder.Decode(&query)
	if err != nil {
		recordResponse.Code = 0
		recordResponse.Message = err.Error()
	}

	ctx := context.Background()

	records, e := s.Application.Queries().GetRecordsQuery.Handle(ctx, &query)
	recordResponse.Records = records

	if e != nil {
		recordResponse.Code = int(e.Code)
		recordResponse.Message = e.Message
	}

	if err := json.NewEncoder(w).Encode(&recordResponse); err != nil {
		return err
	}

	return nil
}

func logRequest(req *http.Request) {
	requestID := req.Header.Get(requestId)
	if requestID == "" {
		requestID = "unknown"
	}
	logger.Println(requestID, req.Method, req.URL.Path, req.RemoteAddr, req.UserAgent())
}
