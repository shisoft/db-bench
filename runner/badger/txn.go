package badger

import (
	"github.com/dgraph-io/badger"
)

type Txn struct {
	txn *badger.Txn
}

func (txn Txn) Set(key []byte, val []byte) error {
	return txn.txn.Set(key, val)
}

func (txn Txn) Read(key []byte) ([]byte, error)  {
	item, err := txn.txn.Get(key)
	if err != nil {
		return nil, err
	}
	return item.Value()
}

func (txn Txn) Seek(key []byte) (interface{}, []byte, error) {
	iter := txn.txn.NewIterator(badger.DefaultIteratorOptions)
	iter.Seek(key)
	item, err := iter.Item().Value()
	return iter, item, err
}

func (txn Txn) Next(cursor interface{}) ([]byte, error) {
	iter := cursor.(*badger.Iterator)
	iter.Next()
	return iter.Item().Value()
}
