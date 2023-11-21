package mcpegsde

import (
	"github.com/midnightfreddie/goleveldb/leveldb"
)

type MCBEWorld struct {
	ldb  *leveldb.DB
	path string
}
