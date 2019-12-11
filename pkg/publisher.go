package pkg

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"
)

// Publisher refer to the domain selling ads
type Publisher struct {
	Sellers []*Seller `json:"sellers"`
}

// Seller contains all the information for a particular seller
type Seller struct {
	PublisherName string `json:"publisherName"`
	Domain        string `json:"domain"`
	AccountID     string `json:"accountId"`
	TypeOfAccount string `json:"typeOfAccount"`
	CertAuthID    string `json:"certAuthId"`
}

func encode(s *Seller) ([]byte, error) {
	var buff bytes.Buffer
	err := gob.NewEncoder(&buff).Encode(s)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func decode(b []byte) (*Seller, error) {
	s := &Seller{}
	err := gob.NewDecoder(bytes.NewBuffer(b)).Decode(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Seller) key() []byte {
	sd := strings.Split(s.Domain, ".")
	return []byte(fmt.Sprintf("ads:%s:%s:%s", s.PublisherName, sd[0], s.AccountID))
}

func adsPrefix(name string) []byte {
	return []byte(fmt.Sprintf("ads:%s", name))
}
