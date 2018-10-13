package runner

type Db interface {
	Init(f string) error
	Halt() error
	Txn(fn func(txn Transaction) error) error
}

type Transaction interface {
	Insert(key []byte, val []byte) error
	Update(key []byte, val []byte) error
	Read(key []byte) ([]byte, error)


	Seek(key []byte) (interface{}, error)
	Next(cursor interface{}) ([]byte, error)
}