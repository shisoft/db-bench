package badger

import (
	"github.com/dgraph-io/badger"
)

type Db struct {
	db *badger.DB
}

func (db Db) Init(f string) error {
	opts := badger.DefaultOptions
	opts.Dir = f
	opts.ValueDir = f
	innerDb, err := badger.Open(opts)
	db.db = innerDb
	return err
}

func (db Db) Halt() error  {
	return db.db.Close()
}

func (db Db) Txn(fn func(txn *Txn) error) error {
	return db.db.Update(func(txn *badger.Txn) error {
		return fn(&Txn{ txn })
	})
}