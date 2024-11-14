package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const PROMPT = ":3000 > "

func main() {
	store := NewStore()

	go func() {
		listener, err := net.Listen("tcp", ":3000")
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}

		defer listener.Close()

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("Failed to accept connection:", err)
				continue
			}
			go handleConnection(conn, store)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Server running")

	for {
    fmt.Printf(PROMPT)
		command := scanner.Scan()
    if !command {
      return
    }

    line := scanner.Text()

		result := HandleCommand(store, line)
		fmt.Println(result)
	}
}

func handleConnection(conn net.Conn, store *Store) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		command := scanner.Text()
		response := HandleCommand(store, command)
		conn.Write([]byte(response + "\n"))
	}
	if err := scanner.Err(); err != nil {
		log.Println("Connection error:", err)
	}
}
