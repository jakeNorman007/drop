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
    if len(parts) != 2 {
      return "GET: gets key and associated value"
    }

    return store.Get(parts[1])

  case "PUTKEY":
    if len(parts) == 1 {
      return "PUTKEY: edits value of key"
    }

    key := parts[1]
    newKey := parts[2]

    exists := store.Get(parts[1])
    fmt.Printf(exists)
    if exists == "" {
      return fmt.Sprintf("Key %v does not exist, create it by running SET command", key)
    }

    store.Set(key, newKey)
    return fmt.Sprintf("Key '%s' updated to '%s'", key, newKey)

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
