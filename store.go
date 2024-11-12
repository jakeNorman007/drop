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

func (s *Store) Set(key, value string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	return "OK"
}

func (s *Store) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, exists := s.data[key]
	if !exists {
		return "(nil)"
	}
	return value
}

func (s *Store) Delete(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.data[key]
	if !exists {
		return "(nil)"
	}
	delete(s.data, key)
	return "OK"
}

func (s *Store) List() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var output string
	for key, value := range s.data {
		output += fmt.Sprintf("%s: %s\n", key, value)
	}
	return output
}
