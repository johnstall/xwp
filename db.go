package xwp

import (
	"github.com/boltdb/bolt"
	// "log"
	// "strings"
	"time"
)

type DB struct {
	db *bolt.DB
}

func Open(path string) (*DB, error) {
	// Open the underlying database.
	d, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	db := &DB{db: d}

	if err := db.db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("xwp"))

		return nil
	}); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
func (db *DB) Close() error {
	return db.db.Close()
}
