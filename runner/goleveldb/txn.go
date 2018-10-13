package goleveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

type Txn struct {
	txn *leveldb.Transaction
	db *leveldb.DB
}

func (txn Txn) Set(key []byte, val []byte) error {
	if txn.txn != nil {
		return txn.txn.Put(key, val, nil)
	} else {
		return txn.db.Put(key, val, nil)
	}
}

func (txn Txn) Read(key []byte) ([]byte, error) {
	if txn.txn != nil {
		return txn.txn.Get(key, nil)
	} else {
		return txn.db.Get(key, nil)
	}
}

func (txn Txn) Seek(key []byte) (interface{}, []byte, error) {
	if txn.txn != nil {
		iter := txn.txn.NewIterator(nil, nil)
		iter.Seek(key)
		return iter, iter.Value(), nil
	} else {
		iter := txn.db.NewIterator(nil, nil)
		iter.Seek(key)
		return iter, iter.Value(), nil
	}
}

func (txn Txn) Next(cursor interface{}) ([]byte, error) {
	iter := cursor.(iterator.Iterator)
	iter.Next()
	return iter.Value(), nil
}