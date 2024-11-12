package main

import (
	"strings"
  "regexp"
)

func HandleCommand(store *Store, command string) string {
  re := regexp.MustCompile(`"([^"]+)"|\S+`)
	parts := re.FindAllString(command, -1)
	if len(parts) == 0 {
		return "Invalid command"
	}

	cmd := strings.ToUpper(parts[0])
	switch cmd {
	case "SET":
		if len(parts) < 3 {
			return "Usage: SET key value"
		}

    value := strings.Join(parts[2:], " ")
    value = strings.Trim(value, "\"")

		return store.Set(parts[1], value)
	case "GET":
		if len(parts) != 2 {
			return "Usage: GET key"
		}
		return store.Get(parts[1])
	case "DELETE":
		if len(parts) != 2 {
			return "Usage: DELETE key"
		}
		return store.Delete(parts[1])
	case "LIST":
		return store.List()
	default:
		return "Unknown command"
	}
}
