package boltdb

import "github.com/boltdb/bolt"

type Txn struct {
	txn *bolt.Tx
}