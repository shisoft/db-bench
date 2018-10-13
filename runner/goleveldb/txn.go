package goleveldb

import "github.com/syndtr/goleveldb/leveldb"

type Txn struct {
	txn *leveldb.Transaction
}