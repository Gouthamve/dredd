package db

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

// Bolter is the boltdb backed implementation
// of Database interface
type Bolter struct {
	db *bolt.DB
}

// NewBolter returns a new Bolter
func NewBolter(path string) (*Bolter, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	return &Bolter{
		db: db,
	}, nil
}

// GetProblem gets the problem with the id
func (b *Bolter) GetProblem(id string) (p Problem, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("problems"))
		if b == nil {
			return nil
		}

		bytedP := b.Get([]byte(id))
		if v == nil {
			return nil
		}

		if err := json.Unmarshal(bytedP, &p); err != nil {
			return err
		}

		return nil
	})

	return
}
