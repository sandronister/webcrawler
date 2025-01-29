package badger

import (
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/badger/v3"
)

func (m *Model) Insert(url, content string) error {
	if m.db == nil {
		return fmt.Errorf("database is not initialized")
	}

	err := m.db.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(url), []byte(content)).WithTTL(24 * time.Hour)
		err := txn.SetEntry(entry)
		if err != nil {
			log.Printf("failed to set entry: %v", err)
		}
		return err
	})

	if err != nil {
		log.Printf("failed to insert data: %v", err)
	}
	return err
}
