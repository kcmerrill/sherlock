package sherlock

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Web starts our webserver on a given port
func (s *Sherlock) Web(port string) {
	r := mux.NewRouter()

	// setup our routes
	r.HandleFunc("/{entity}", s.WebEntity)
	r.HandleFunc("/{entity}/{property}/{action}", s.WebProcess)
	r.HandleFunc("/{entity}/{event}", s.WebEvent)

	// set some defaults
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// start serving
	log.Fatal(srv.ListenAndServe())
}

// WebEntity handler for web requests with new messages
func (s *Sherlock) WebEntity(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	j, err := s.E(vars["entity"]).String()
	if err == nil {
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, j)
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "{}")
	}
}

// WebEvent handler for web events
func (s *Sherlock) WebEvent(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	s.E(vars["entity"]).Event(vars["event"])
	j, err := s.E(vars["entity"]).String()
	if err == nil {
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, j)
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "{}")
	}
}

// WebProcess handler for web requests with new messages
func (s *Sherlock) WebProcess(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	s.Process(vars["entity"]+"|"+vars["property"]+"|"+vars["action"], "|")

	j, err := s.E(vars["entity"]).String()
	if err == nil {
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, j)
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "{}")
	}
}
