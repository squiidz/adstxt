package pkg

import (
	"log"
)

// AdsServer is a wrapper for the AdsStore
type AdsServer struct {
	store *adsStore
}

// NewAdsServer returns an AdsServer reference
func NewAdsServer(dbpath string) *AdsServer {
	s := &AdsServer{store: newStore(dbpath)}
	if err := s.store.Open(); err != nil {
		// Can't do anything if the store fail
		log.Fatal(err)
		return nil
	}
	return s
}

// GetPublisher return a Publisher containing the sellers entries
func (s *AdsServer) GetPublisher(name string) (*Publisher, error) {
	p, err := s.store.getPublisher(name)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// AddPublisher insert the publisher ads.txt file
// content in the underlying AdsServer database
func (s *AdsServer) AddPublisher(p *Publisher) error {
	return s.store.addPublisher(p)
}
