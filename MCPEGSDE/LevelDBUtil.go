package mcpegsde

import (
	"errors"
	"os"

	"github.com/midnightfreddie/goleveldb/leveldb"
)

type MCBEWorld struct {
	Ldb  *leveldb.DB
	Path string
}

// OpenMCBEWorld opens a Minecraft Bedrock Edition world at the
// specified path. Then returns a pointer to a MCBEWorld struct
// containing the LevelDB database and the path to the world.
func OpenMCBEWorld(Path string) (*MCBEWorld, error) {
	world := MCBEWorld{nil, Path}
	var dbPath = Path + "/db"

	fileInfo, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		return &world, errors.New("LevelDB database does not exist. This must be ran on a valid MCBE world directory.")
	}
	if err != nil {
		return &world, err
	}
	if !fileInfo.IsDir() {
		return &world, errors.New("LevelDB database does not exist. This must be ran on a valid MCBE world directory.")
	}

	world.Ldb, err = leveldb.OpenFile(dbPath, nil)
	if err != nil {
		if world.Ldb != nil {
			_ = world.Ldb.Close()
		}
		return &world, err
	}
	return &world, nil
}

// Close closes the LevelDB database.
func (w *MCBEWorld) Close() error {
	return w.Ldb.Close()
}

// FilePath returns the path to the Minecraft Bedrock Edition world.
func (w *MCBEWorld) FilePath() string {
	return w.Path
}

// GetKeys returns all keys in the LevelDB database as a [][]byte
// slice. This is a memory-intensive operation and should be used
// with caution.
func (w *MCBEWorld) GetKeys() ([][]byte, error) {
	keylist := [][]byte{}
	iter := w.Ldb.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		tmp := make([]byte, len(key))
		copy(tmp, key)
		keylist = append(keylist, tmp)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return keylist, err
	}
	return keylist, nil
}

// GetFromKey returns the value from the LevelDB database at the
// specified key. This is a convenience function that wraps the
// LevelDB Get function. This function copies the data from the
// LevelDB database, so the returned []byte slice is a copy of the
// LevelDB database. This means that the returned []byte slice can be
// modified. Then use the PutFromKey function to write the modified
// data back to the LevelDB database.
func (w *MCBEWorld) GetFromKey(key []byte) ([]byte, error) {
	temp, err := w.Ldb.Get(key, nil)

	val := make([]byte, len(temp))
	copy(val, temp)
	return val, err
}

// GetFromKeyUnsafe returns the value from the LevelDB database at the
// specified key. This is a convenience function that wraps the
// LevelDB Get function. This function does not copy the data from the
// LevelDB database, so the returned []byte slice is a reference to the
// LevelDB database. This means that the returned []byte slice should
// not be modified.
// Can error if the LevelDB Get function errors.
func (w *MCBEWorld) GetFromKeyUnsafe(key []byte) ([]byte, error) {
	return w.Ldb.Get(key, nil)
}

// PutFromKey puts a value into the LevelDB database at the specified
// key. This is a convenience function that wraps the LevelDB Put
// function.
// Can error if the LevelDB Put function errors.
func (w *MCBEWorld) PutFromKey(key []byte, value []byte) error {
	return w.Ldb.Put(key, value, nil)
}
