package kvstore

import (
    "sync"
    "time"
)

type Data struct {
    value      string
    expiration time.Time
}

type KeyValueStore struct {
    mu   sync.RWMutex
    data map[string]Data
}

func NewKeyValueStore() *KeyValueStore {
    return &KeyValueStore{
        data: make(map[string]Data),
    }
}
