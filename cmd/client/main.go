package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/DanielTitkov/sillyquotes/internal/pow"
)

func request(address string) error {
	tcpServer, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpServer)
	if err != nil {
		return err
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}

	challenge := strings.TrimSpace(string(buffer[:n]))

	log.Printf("Recieved challenge: '%s' of len %d", challenge, len(challenge))

	solution := pow.SolveChallenge(fmt.Sprint(challenge), pow.DefaultRequiredZeros)

	log.Printf("Generated solution: '%s'", solution)
	_, err = conn.Write([]byte(solution))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		received := scanner.Text()
		log.Println("Received message:", received)
	}
	if err := scanner.Err(); err != nil {
		log.Println("Read data failed:", err.Error())
	}

	return nil
}

func main() {
	address := os.Getenv("ADDRESS")
	if address == "" {
		address = "localhost:8099"
	}

	for {
		if err := request(address); err != nil {
			log.Printf("Request failed with error: %s", err)
		}
		time.Sleep(time.Second * 3)

	}
}
