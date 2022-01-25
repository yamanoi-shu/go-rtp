package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("localhost"),
		Port: 8080,
	}

	conn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fmt.Println("listening.....")

	var buf = make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("receive from address: ", addr)
		go handleUDP(buf[:n])
		buf = make([]byte, 1024)
	}
}

func handleUDP(buf []byte) {

	fmt.Println(string(buf))
}
