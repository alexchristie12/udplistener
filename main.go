package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)

// Create the UDP listener server
func startUDPServer(port string) (*net.UDPConn, error) {
	listenAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	conn, err := startUDPServer(port)
	if err != nil {
		log.Fatal("Could not start UDP listener server: ", err)
	}
	defer conn.Close()

	log.Printf("UDP server listening on: udp://127.0.0.1%s\n", port)

	// Start listening for messages
	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Could not read UDP message: ", err)
			continue
		}
		log.Printf("Received message from %s:\n%s\n", addr, buffer[:n])
	}
}
