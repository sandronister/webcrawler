package badger

import (
	"time"

	"github.com/dgraph-io/badger/v3"
)

func (m *Model) Insert(url, content string) error {
	err := m.db.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(url), []byte(content)).WithTTL(24 * time.Hour)
		err := txn.SetEntry(entry)
		return err
	})
	return err
}
