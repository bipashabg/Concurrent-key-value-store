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

//op set
func (kv *KeyValueStore) Set(key, val string, ttl time.Duration) {
    kv.mu.Lock()
    defer kv.mu.Unlock()

    kv.data[key] = Data{
        value:      val,
        expiration: time.Now().Add(ttl),
    }
}

//op get
func (kv *KeyValueStore) Get(key string) (string, bool) {
    kv.mu.Lock()
    defer kv.mu.Unlock()
    val, ok := kv.data[key]

    // Check for Item Expiry
    if val.expiration.IsZero() || time.Now().Before(val.expiration) {
        return val.value, ok
    }

    // Delete if Expired
    delete(kv.data, key)
    return "", false
}

//op del
func (kv *KeyValueStore) Delete(key string) bool {
    kv.mu.Lock()
    defer kv.mu.Unlock()

    _, ok := kv.data[key]
    if !ok {
        return false
    }
    delete(kv.data, key)
    return true
}
