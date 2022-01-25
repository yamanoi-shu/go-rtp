package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	conn.Write([]byte("hello"))
}
