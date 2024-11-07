package main

import (
  "fmt"
  "reflect"
  "testing"
)

func TestStore(t *testing.T) {
  db := New() 

  db.Set("key1", "value1")
  if value, ok := db.Get("key1"); !ok || value != "value1" {
    t.Errorf("Expected value1, got %v", value)
  }

  db.Set("key2", 32)
  if value, ok := db.Get("key2"); !ok || value != 32 {
    t.Errorf("Expected 32, got %v", value)
  }

  db.Get("key1")
  if value, ok := db.Get("key1"); !ok || value != "value1" {
    t.Errorf("Expected value1, got %v", value)
  }

  value, exists := db.Get("key1")
  if !exists {
    t.Fatalf("key1 does not exist")
  }

  value2, exists := db.Get("key2")
  if !exists {
    t.Fatalf("key2 does not exist")
  }

  listData := db.List()
  if !reflect.DeepEqual(listData, db.store) {
    t.Errorf("You have no data in your drop store")
  }

  fmt.Printf("key1's value is: %v\n", value)
  fmt.Printf("key2's value is: %v\n", value2)

  db.Delete("key1")
  if _, ok := db.Get("key1"); ok {
    t.Errorf("Expected key1 to be deleted")
  }
}
