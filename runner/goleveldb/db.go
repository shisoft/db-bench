package goleveldb

import (
	"github.com/shisoft/db-bench/runner"
	"github.com/syndtr/goleveldb/leveldb"
)

type Db struct {
	db *leveldb.DB
}

func (db Db) Init(f string) error {
	inner, err := leveldb.OpenFile(f, nil)
	db.db = inner
	return err
}

func (db Db) Halt() error  {
	return db.db.Close()
}

func (db Db) Update(fn func(txn runner.Transaction) error) error  {
	txn, err := db.db.OpenTransaction()
	if err != nil { return err }
	itxn := Txn{ txn, nil }
	return fn(itxn)
}

func (db Db) View(fn func(txn runner.Transaction) error) error {
	itxn := Txn{ nil, db.db }
	return fn(itxn)
}