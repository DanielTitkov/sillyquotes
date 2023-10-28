package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strconv"
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

	buffer := make([]byte, pow.DefaultChallengeLength)
	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}

	challenge := buffer[:n]

	log.Printf("Received challenge: '%s' of len %d", string(challenge), len(challenge))

	solution, err := pow.SolveChallenge(challenge, pow.DefaultRequiredZeros)
	if err != nil {
		return err
	}

	log.Printf("Generated solution: '%d'", solution)
	_, err = conn.Write([]byte(strconv.FormatUint(uint64(solution), 10))) // this is intentionally passed as a string
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
