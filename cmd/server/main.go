package main

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/DanielTitkov/sillyquotes/internal/handler"
	"github.com/DanielTitkov/sillyquotes/internal/quote"
)

var (
	listen       string
	readTimeout  int
	writeTimeout int
)

func init() {
	flag.StringVar(&listen, "listen", "localhost:8099", "Server address")
	flag.IntVar(&readTimeout, "readTimeout", 600, "Read timeout in milliseconds")
	flag.IntVar(&writeTimeout, "writeTimeout", 150, "Write timeout in milliseconds")
	flag.Parse()
}

func main() {
	h := handler.New(
		quote.NewProvider(),                          // quote provider
		time.Duration(readTimeout)*time.Millisecond,  // read timeout
		time.Duration(writeTimeout)*time.Millisecond, // write timeout
	)

	log.Printf("Server starting, listening at %s, timeouts: read - %dms, write - %dms", listen, readTimeout, writeTimeout)
	listen, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}

	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalf("Error getting connection: %s", err)
		}
		go h.HandlePoW(conn)
	}
}
