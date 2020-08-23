package dictionary

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger"
)

// Dictionary type to instanciate a DB connection
type Dictionary struct {
	db *badger.DB
}

// Entry manage our keyboard input
type Entry struct {
	Word       string
	Definition string
	CreatedAt  time.Time
}

func (e Entry) String() string {
	created := e.CreatedAt.Format(time.Stamp)
	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, created)
}

// New generates a new connection with options for badger
func New(dir string) (*Dictionary, error) {
	opts := badger.DefaultOptions("./badger")
	opts.Dir = dir
	opts.ValueDir = dir

	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	dict := &Dictionary{
		db: db,
	}
	return dict, nil
}

// Close like mentionned just close the badger db connection
func (d *Dictionary) Close() {
	d.db.Close()
}
