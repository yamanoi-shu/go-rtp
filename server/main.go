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

	p := packet.NewRTPPacket()
	buf = p.Header.Unmarshal(buf)
	buf = p.Payload.Depacketize(buf)
	fmt.Println(p.Header.String())
	fmt.Println(p.Payload.String())
	fmt.Println()
}
