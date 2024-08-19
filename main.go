package main

import (
	"errors"
	"log"
	"net"
	"os"
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

func getIPAddr() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() || ipnet.IP.IsLinkLocalUnicast() {
			continue
		}
		ip := ipnet.IP.To4()
		if ip != nil && ip[0] == 192 && ip[1] == 168 {
			return ip, nil
		}
	}
	return nil, errors.New("could not get ip address")
}

func main() {
	port := ":1234"
	conn, err := startUDPServer(port)
	if err != nil {
		log.Fatal("Could not start UDP listener server: ", err)
	}
	defer conn.Close()
	ip, err := getIPAddr()
	if err != nil {
		log.Fatal("Could not get IP address: ", err)
	}
	log.Printf("UDP server listening on: %v%s\n", ip, port)

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
