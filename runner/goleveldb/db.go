package goleveldb

import "github.com/syndtr/goleveldb/leveldb"

type Db struct {
	db *leveldb.DB
}