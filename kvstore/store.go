package kvstore

import (
	"errors"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

// Store implements a key-value store. May be backed by a database like LevelDB.
type Store interface {

	// Put stores the given key / value pair in the database, overwriting any existing value for that key.
	Put(key []byte, value []byte) error

	// Get retrieves the value for the given key. Returns a ErrNotFound error if the key does not exist.
	Get(key []byte) ([]byte, error)

	// Delete removes the key from the database. Does not return an error if the key does not exist.
	Delete(key []byte) error

	// DeleteBatch atomically removes a list of keys from the database.
	DeleteBatch(keys [][]byte) error

	// WriteBatch atomically writes a list of key / value pairs to the database. The key at index i in the keys slice
	// corresponds to the value at index i in the values slice.
	WriteBatch(keys, values [][]byte) error

	// NewIterator returns an iterator that can be used to iterate over a subset of the keys in the database.
	// Only keys with the given prefix will be iterated.
	// TODO describe snapshot behavior
	NewIterator(prefix []byte) (iterator.Iterator, error)

	// Shutdown shuts down the store, flushing any remaining data to disk.
	Shutdown() error

	// Destroy shuts down and permanently deletes all data in the store.
	Destroy() error
}

// ErrNotFound is returned when a key is not found in the database.
var ErrNotFound = errors.New("not found")