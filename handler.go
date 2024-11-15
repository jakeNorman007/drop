package main

import (
  "fmt"
  "regexp"
  "strings"
)

func HandleCommand(store *Store, command string) string {
  regex := regexp.MustCompile(`"([^"]+)"|\S+`)

  parts := regex.FindAllString(command, -1)
  if len(parts) == 0 {
    return "Invalid command"
  }

  cmd := strings.ToUpper(parts[0])

  switch cmd {
  case "SET":
    if len(parts) < 3 {
      return "SET: sets key and value"
    }

    value := strings.Join(parts[2:], " ")
    value = strings.Trim(value, "\"")

    return store.Set(parts[1], value)

  case "GET":
    if len(parts) < 1 {
      return "GET: gets key and associated value"
    }

    key := parts[1]
    fmt.Println("Looking for key:", key)

    value, exists := store.Get(key)
    if !exists {
      return "nil"
    }

    return value

  case "PUTKEY":
    if len(parts) < 2 {
      return "Usage: EDIT key new_value"
    }

    key := parts[1]
    newValue := parts[2]

    value, exists := store.Get(key)
    if !exists {
      return fmt.Sprintf("Key '%s' does not exist. Use SET to create it.", key)
    }

    store.Set(newValue, value)
    store.Delete(key)

    return fmt.Sprintf("Key '%s' updated to '%s'", key, newValue)

  case "DEL":
    if len(parts) != 2 {
      return "DEL: deletes key and associated values"
    }

    return store.Delete(parts[1])

  case "LIST":
    return store.List()

  default:
    return "Unknown command"
  }
}
