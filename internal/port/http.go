package port

import (
	"context"
	"encoding/json"
	"github.com/ybalcin/storage-api/internal/application"
	"github.com/ybalcin/storage-api/internal/application/commandhandler"
	"github.com/ybalcin/storage-api/internal/application/queryhandler"
	"log"
	"net/http"
	"os"
)

type (
	httpServer struct {
		Application *application.Application
	}

	Handler struct {
		H      func(rw http.ResponseWriter, req *http.Request) error
		Method string
	}
)

const (
	contentType     = "Content-Type"
	applicationJson = "application/json"
	requestId       = "X-Request-ID"

	RequiredKey = "key is required"
)

var (
	logger = log.New(os.Stdout, "in: ", log.LstdFlags)
)

// NewHttpServer initializes a new in server input port
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

	if req.Method != h.Method {
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

	//if err != nil {
	//	switch e := err.(type) {
	//	case common.Error:
	//		logger.Printf("HTTP status: %d - Message: %s", e.Status(), e.Error())
	//		http.Error(w, "", e.Status())
	//	default:
	//		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	//	}
	//}
}

// AddCacheEntryHandler adds cache entry to in memory
func (s *httpServer) AddCacheEntryHandler(w http.ResponseWriter, req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	var command commandhandler.AddCacheEntryCommand

	if err := decoder.Decode(&command); err != nil {
		return err
	}

	err := s.Application.Commands().AddCacheEntryCommand.Handle(&command)
	if err != nil {
		// return
	}

	return nil
}

// GetCacheEntryHandler gets entry from cache
func (s *httpServer) GetCacheEntryHandler(w http.ResponseWriter, req *http.Request) error {
	key := req.URL.Query().Get("key")

	queryResult, err := s.Application.Queries().GetCacheEntryQuery.Handle(&queryhandler.GetCacheEntryQuery{Key: key})
	if err != nil {
		// return
	}

	if err := json.NewEncoder(w).Encode(queryResult); err != nil {
		return err
	}
	return nil
}

// GetRecordsHandler gets records
func (s *httpServer) GetRecordsHandler(w http.ResponseWriter, req *http.Request) error {
	decoder := json.NewDecoder(req.Body)

	var query queryhandler.GetRecordsQuery

	if err := decoder.Decode(&query); err != nil {
		return err
	}

	ctx := context.Background()

	records, err := s.Application.Queries().GetRecordsQuery.Handle(ctx, &query)

	var code int
	var message string
	if err == nil {
		code = 0
		message = "success"
	} else {
		code = err.Int()
		message = err.Message
	}

	recordResponse := RecordHttpResponse{
		Code:    code,
		Message: message,
		Records: records,
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
