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

    value, exists := store.Get(key)
    if !exists {
      return "nil"
    }

    return value

  case "EKEY":
    if len(parts) < 2 {
      return "EKEY: edit the key's value"
    }

    key := parts[1]
    newKey := parts[2]

    value, exists := store.Get(key)
    if !exists {
      return fmt.Sprintf("Key '%s' does not exist. Use SET to create.", key)
    }

    store.Set(newKey, value)
    store.Delete(key)

    return fmt.Sprintf("Key '%s' updated to '%s'", key, newKey)

  case "EVALUE":
    if len(parts) < 2 {
      return "EVALUE: edit the value of a selected key"
    }

    key := parts[1]
    newKey := parts[1]
    newValue := parts[2]

    _, exists := store.Get(key)
    if !exists {
      return fmt.Sprintf("Key '%s' does not exist. Use SET to create.", key)
    }

    store.Delete(key)
    store.Set(newKey, newValue)

    return fmt.Sprintf("Key '%s's'value updated to '%s'", key, newValue)

  case "DEL":
    if len(parts) != 2 {
      return "DEL: deletes key and associated values"
    }

    return store.Delete(parts[1])

  case "LIST":
    return store.List()

  default:
    return "Command is unknown"
  }
}
