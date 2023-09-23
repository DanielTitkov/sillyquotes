package main

import (
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/DanielTitkov/sillyquotes/internal/pow"
	"github.com/DanielTitkov/sillyquotes/internal/quote"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	challenge := pow.GenerateChallenge(pow.DefaultChallengeLength)
	log.Printf("New connection. Sending challenge: '%s'", challenge)
	start := time.Now()

	_, err := conn.Write([]byte(challenge))
	if err != nil {
		log.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	solution := strings.TrimSpace(string(buffer[:n]))

	log.Printf("Got solution '%s' for the challenge '%s'", solution, challenge)

	if pow.CheckPoW(challenge, solution, pow.DefaultRequiredZeros) {
		log.Printf("Access granted in %s", time.Since(start))
		conn.Write([]byte(quote.Random()))
	} else {
		log.Printf("Access denied in %s", time.Since(start))
		conn.Write([]byte("Incorrect PoW solution."))
	}
}

func main() {
	address := os.Getenv("LISTEN")
	if address == "" {
		address = "localhost:8099"
	}

	log.Printf("Server starting, listening at %s", address)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}

	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalf("Error getting connection: %s", err)
		}
		go handleConnection(conn)
	}
}
