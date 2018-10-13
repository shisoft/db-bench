package boltdb

import (
	"github.com/boltdb/bolt"
)

type Db struct {
	db *bolt.DB
}