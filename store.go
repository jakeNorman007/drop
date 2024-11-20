package main

import (
  "fmt"
  "sync"
)

type Store struct {
  mu    sync.RWMutex
  data  map[string]string
}

func NewStore() *Store {
  return &Store {
    data: make(map[string]string),
  }
}

type GenFunctions interface {
  Set(key, value string) string
  Get(key string) (string, bool)
  Delete(key string)
  List()
}

func (s *Store) Set(key, value string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

	return "SET confirmed"
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, exists := s.data[key]

	return value, exists
}

func (s *Store) Delete(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.data[key]
	if !exists {
		return "no values to DELETE"
	}

	delete(s.data, key)

	return "DELETE confirmed"
}

func (s *Store) List() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var output string

	for key, value := range s.data {
    output += fmt.Sprintf("key: %s, value: %s\n", key, value)
	}

	return output
}
