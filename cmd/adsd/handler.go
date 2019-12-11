package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/squiidz/adstxt/pkg"

	"github.com/go-zoo/bone"
)

type Server struct {
	*pkg.AdsServer
}

func NewServer(dbpath string) *Server {
	return &Server{pkg.NewAdsServer(dbpath)}
}

func (s *Server) getPublisher(rw http.ResponseWriter, req *http.Request) {
	val := bone.GetValue(req, "publisher")
	name := strings.ToLower(val)
	publisher, err := s.GetPublisher(name)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(publisher)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type AddPublisherReq struct {
	Domain string `json:"domain"`
}

func (s *Server) addPublisher(rw http.ResponseWriter, req *http.Request) {
	apr := &AddPublisherReq{}
	err := json.NewDecoder(req.Body).Decode(apr)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := pkg.ProcessDomain(apr.Domain)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	if err = s.AddPublisher(p); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
