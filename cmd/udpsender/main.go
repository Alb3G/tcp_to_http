package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const NETWORK = "udp"
const ADDR = "localhost:42069"

func main() {
	udpAddr, err := net.ResolveUDPAddr(NETWORK, ADDR)
	if err != nil {
		log.Fatalf("Error initializating udp sender: %s", err)
	}

	udpConn, err := net.DialUDP(NETWORK, nil, udpAddr)
	if err != nil {
		log.Fatalf("Error connection at udp sender: %s", err)
	}
	defer udpConn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error connection at udp sender: %s", err)
		}

		udpConn.Write([]byte(line))
	}
}
