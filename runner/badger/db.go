package badger

import (
	"github.com/dgraph-io/badger"
	"github.com/shisoft/db-bench/runner"
)

type Db struct {
	db *badger.DB
}

func (db Db) Init(f string) error {
	opts := badger.DefaultOptions
	opts.Dir = f
	opts.ValueDir = f
	inner, err := badger.Open(opts)
	db.db = inner
	return err
}

func (db Db) Halt() error  {
	return db.db.Close()
}

func (db Db) Update(fn func(txn runner.Transaction) error) error {
	return db.db.Update(func(txn *badger.Txn) error {
		return fn(&Txn{ txn })
	})
}

func (db Db) View(fn func(txn runner.Transaction) error) error {
	return db.db.View(func(txn *badger.Txn) error {
		return fn(&Txn{ txn })
	})
}