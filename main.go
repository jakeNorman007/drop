package main

import (
  "fmt"
  "sync"
)

type Store struct {
  store map[string]interface{}
  mu sync.RWMutex
}

func New() *Store {
  return &Store {
    store: make(map[string]interface{}),
  }
}

func (db *Store) Set(key string, value interface{}) {
  db.mu.Lock()
  defer db.mu.Unlock()
  db.store[key] = value
}

func (db *Store) Get(key string) (interface{}, bool) {
  db.mu.RLock()
  defer db.mu.RUnlock()
  value, exists := db.store[key]
  return value, exists
}

func (db *Store) Delete(key string) {
  db.mu.Lock()
  defer db.mu.Unlock()
  delete(db.store, key)
}

func (db *Store) Clear() {
  db.mu.Lock()
  defer db.mu.Unlock()
  db.store = make(map[string]interface{})
}

func (db *Store) List() map[string]interface{} {
  db.mu.RLock()
  defer db.mu.RUnlock()

  list := make(map[string]interface{})

  for key, value := range db.store {
    list[key] = value
    fmt.Printf("Key: %s, Value: %v\n", key, value)
  }

  return list
}

type StoreFunctions interface {
  Set(key string, value interface{})
  Get(key string) (interface{}, bool)
  Delete(key string)
  Clear()
  List()
}
