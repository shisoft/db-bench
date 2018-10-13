package boltdb

import "github.com/boltdb/bolt"

type Txn struct {
	txn *bolt.Tx
	bucketName []byte
}

func (txn Txn) Set(key []byte, val []byte) error  {
	b := txn.txn.Bucket(txn.bucketName)
	return b.Put(key, val)
}

func (txn Txn) Read(key []byte) ([]byte, error) {
	b := txn.txn.Bucket(txn.bucketName)
	return b.Get(key), nil
}

func (txn Txn) Seek(key []byte) (interface{}, []byte, error) {
	b := txn.txn.Bucket(txn.bucketName)
	c := b.Cursor()
	_, val := c.Seek(key)
	return c, val, nil
}

func (txn Txn) Next(cursor interface{}) ([]byte, error) {
	c := cursor.(*bolt.Cursor)
	_, val := c.Next()
	return val, nil
}