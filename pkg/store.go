package pkg

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

type storeState int

const (
	openStore storeState = iota
	closeStore
)

var (
	// ErrStoreClose is used when a method is called on a closed adsStore reference
	ErrStoreClose = fmt.Errorf("you need to open the store before using it")
)

// adsStore is used to store all the publishers and sellers
type adsStore struct {
	state  storeState
	dbpath string
	db     *badger.DB
}

// newStore returns a adsStore reference
func newStore(dbpath string) *adsStore {
	return &adsStore{
		state:  closeStore,
		dbpath: dbpath,
		db:     nil,
	}
}

// Open is used to initialized the adsStore
// this pattern is used since by convention a "NewX" in go shouldn't return a error
func (s *adsStore) Open() error {
	opts := badger.DefaultOptions(s.dbpath)
	opts.Logger = nil
	db, err := badger.Open(opts)
	if err != nil {
		return err
	}
	s.db = db
	s.state = openStore
	return nil
}

func (s *adsStore) getPublisher(name string) (*Publisher, error) {
	if s.state == closeStore {
		return nil, ErrStoreClose
	}
	p := &Publisher{}
	txn := s.db.NewTransaction(false)
	itr := txn.NewIterator(badger.DefaultIteratorOptions)
	prefix := adsPrefix(name)

	itr.Seek(prefix)
	for itr.ValidForPrefix(prefix) {
		err := itr.Item().Value(func(val []byte) error {
			s, err := decode(val)
			if err != nil {
				return err
			}
			p.Sellers = append(p.Sellers, s)
			return nil
		})
		if err != nil {
			log.Println(err)
			continue
		}
		itr.Next()
	}
	return p, nil
}

func (s *adsStore) addPublisher(p *Publisher) error {
	if s.state == closeStore {
		return ErrStoreClose
	}
	batch := s.db.NewWriteBatch()
	for _, s := range p.Sellers {
		v, err := encode(s)
		if err != nil {
			log.Println(err)
			continue
		}
		err = batch.Set(s.key(), v)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return batch.Flush()
}
