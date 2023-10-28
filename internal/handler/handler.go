package handler

import (
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/DanielTitkov/sillyquotes/internal/pow"
	"github.com/DanielTitkov/sillyquotes/internal/quote"
)

type Handler struct {
	quote        *quote.Provider
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func New(
	quoteProvider *quote.Provider,
	readTimeout time.Duration,
	writeTimeout time.Duration,
) *Handler {
	return &Handler{
		quote:        quoteProvider,
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
	}
}

func (h *Handler) HandlePoW(conn net.Conn) {
	defer conn.Close()
	err := conn.SetReadDeadline(time.Now().Add(h.readTimeout))
	if err != nil {
		log.Println(err)
		return
	}

	err = conn.SetWriteDeadline(time.Now().Add(h.writeTimeout))
	if err != nil {
		log.Println(err)
		return
	}

	challenge := pow.GenerateChallenge(pow.DefaultChallengeLength)
	log.Printf("New connection. Sending challenge: '%s'", challenge)
	start := time.Now()

	_, wErr := conn.Write([]byte(challenge))
	if err != nil {
		log.Printf("Failed to write challenge %v", wErr)
		return
	}

	buffer := make([]byte, 10) // Max length of string representation of uint 32
	// this can be optimised and be passed as an uint32 in 4 bytes
	// but is was intentionally kept as a string for now for simplicity of logging
	// FIXME
	n, _ := conn.Read(buffer)
	solutionStr := strings.TrimSpace(string(buffer[:n]))

	log.Printf("Got solution '%s' for the challenge '%s'", solutionStr, challenge)

	solution, err := strconv.ParseUint(solutionStr, 10, 32)
	if err != nil {
		log.Printf("Access denied in %s", time.Since(start))
		_, wErr = conn.Write([]byte("PoW solution must be integer"))
		if wErr != nil {
			log.Printf("Failed to write response %v", wErr)
			return
		}
	}

	if pow.CheckPoW(challenge, uint32(solution), pow.DefaultRequiredZeros) {
		log.Printf("Access granted in %s", time.Since(start))
		_, wErr = conn.Write([]byte(h.quote.Random()))
	} else {
		log.Printf("Access denied in %s", time.Since(start))
		_, wErr = conn.Write([]byte("Incorrect PoW solution."))
	}

	if wErr != nil {
		log.Printf("Failed to write response %v", wErr)
		return
	}
}
