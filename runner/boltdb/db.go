package boltdb

import (
	"github.com/boltdb/bolt"
	"github.com/shisoft/db-bench/runner"
)

type Db struct {
	db *bolt.DB
}

const defaultBucket  = "Test"

func (db Db) Init(f string) error {
	inner, err := bolt.Open(f, 0600, nil)
	db.db = inner
	return err
}

func (db Db) Halt() error  {
	return db.db.Close()
}

func (db Db) View(fn func(txn runner.Transaction) error) error {
	return db.db.View(func(tx *bolt.Tx) error {
		return fn(&Txn{ tx, []byte(defaultBucket) })
	})
}

func (db Db) Update(fn func(txn runner.Transaction) error) error {
	return db.db.Update(func(tx *bolt.Tx) error {
		return fn(&Txn{ tx, []byte(defaultBucket) })
	})
}