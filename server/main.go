package main

import (
	"fmt"
	"go-rtp/packet"
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

	var clientCnt int

	var buf = make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if clientCnt == 0 {
			fmt.Println("receive from address: ", addr)
		}
		clientCnt++

		if err != nil {
			log.Fatal(err)
		}

		go handleUDP(buf[:n])
		buf = make([]byte, 1024)
	}
}

func handleUDP(buf []byte) {

	h := &packet.RTPHeader{}
	h.Unmarshal(buf)
	fmt.Println(h.String())
	fmt.Println()
}
