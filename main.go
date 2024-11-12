package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	store := NewStore()

	// Start TCP server
	go func() {
		listener, err := net.Listen("tcp", ":6379")
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
		defer listener.Close()
		fmt.Println("Server is running on :6379")
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("Failed to accept connection:", err)
				continue
			}
			go handleConnection(conn, store)
		}
	}()

	// CLI interface
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("CLI ready. Type commands (SET, GET, DELETE, LIST).")

	for scanner.Scan() {
		command := scanner.Text()
		result := HandleCommand(store, command)
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error reading CLI input:", err)
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
