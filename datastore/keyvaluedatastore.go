package datastore

import (
	"sync"

	"github.com/geekible/memstore/constants"
	"github.com/geekible/memstore/models"
)

type keyValueRecord map[string]any

type KeyValueDataStore struct {
	records keyValueRecord
}

var (
	mu = new(sync.RWMutex)
)

func NewKeyValueDataStore() *KeyValueDataStore {
	return &KeyValueDataStore{
		records: make(keyValueRecord),
	}
}

func (kvs *KeyValueDataStore) Get(key string) (models.KeyValue, error) {
	mu.RLock()
	defer mu.RUnlock()

	value, ok := kvs.records[key]
	if !ok {
		return models.KeyValue{}, constants.ErrNotFound
	}

	kv := models.NewKeyValue().
		SetKey(key).
		SetValue(value).
		Build()

	return *kv, nil
}

func (kvs *KeyValueDataStore) Set(kv models.KeyValue) error {
	mu.Lock()
	defer mu.Unlock()

	kvs.records[kv.GetKey()] = kv.GetValue()
	return nil
}

func (kvs *KeyValueDataStore) Delete(key string) error {
	mu.Lock()
	defer mu.Unlock()

	delete(kvs.records, key)
	return nil
}
