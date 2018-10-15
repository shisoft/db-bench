package runner

import (
	"encoding/binary"
	"github.com/shisoft/db-bench/runner/badger"
	"github.com/shisoft/db-bench/runner/boltdb"
	"github.com/shisoft/db-bench/runner/goleveldb"
)

func InstancateDb(db string) Db  {
	switch db {
	case "leveldb":
		return goleveldb.Db{}
	case "boltdb":
		return boltdb.Db{}
	case "badger":
		return badger.Db{}
	default:
		panic("unknown db")
	}
}

func InitDb(db Db, f string) error {
	return db.Init(f)
}

func HaltDb(db Db) error {
	return db.Halt()
}

func intToBytes(num uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, num)
	return bs
}

func makeContent(num uint32) []byte  {
	bs := make([]byte, 1024)
	binary.BigEndian.PutUint32(bs, num)
	return bs
}

func BulkSet(db Db, num uint32) error {
	return db.Update(func(txn Transaction) error {
		for i := uint32(0); i < num; i++ {
			data := intToBytes(i)
			txn.Set(data, makeContent(num))
		}
		return nil
	})
}

func BulkGet(db Db, num uint32) error {
	chans := make([]chan error, num)
	for i := uint32(0); i < num; i++ {
		chans[i] = make(chan error)
	}
	for i := uint32(0); i < num; i++ {
		go func() {
			chans[i] <- db.View(func(txn Transaction) error {
				_, err := txn.Read(intToBytes(i))
				return err
			})
		}()
	}
	for i := uint32(0); i < num; i++ {
		err := <- chans[i]
		if err != nil {
			return err
		}
	}
	return nil
}

func BulkScan(db Db, num uint32) error  {
	chans := make([]chan error, num)
	for i := uint32(0); i < num; i++ {
		chans[i] = make(chan error)
	}
	for i := uint32(0); i < num; i++ {
		go func() {
			chans[i] <- db.View(func(txn Transaction) error {
				cursor, _, err := txn.Seek(intToBytes(i))
				if err != nil {
					return err
				}
				for j := i; j < num; j++ {
					_, err := txn.Next(cursor)
					if err != nil {
						return err
					}
				}
				return nil
			})
		}()
	}
	for i := uint32(0); i < num; i++ {
		err := <- chans[i]
		if err != nil {
			return err
		}
	}
	return nil
}