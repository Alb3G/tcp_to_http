package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	chanStr := make(chan string)

	var currentLine string

	go func() {
		defer f.Close()
		defer close(chanStr)
		for {
			bytes := make([]byte, 8)
			n, err := f.Read(bytes)
			if err != nil {
				if err == io.EOF {
					if currentLine != "" {
						chanStr <- currentLine
						currentLine = ""
					}
					return
				} else {
					log.Fatalf("Exiting programm with error :%v", err)
				}
			}

			data := string(bytes[:n])

			parts := strings.Split(data, "\n")

			for i := 0; i < len(parts)-1; i++ {
				chanStr <- currentLine + parts[i]
				currentLine = ""
			}

			currentLine += parts[len(parts)-1]
		}
	}()

	return chanStr
}

const PORT = ":42069"

func main() {
	tcp, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Error while listening TCP: %s", err)
	}
	defer tcp.Close()

	for {
		conn, err := tcp.Accept()
		if err != nil {
			log.Fatalf("Error listening tcp connections: %s", err)
		}

		fmt.Println("Connection accepted, processing...")

		for line := range getLinesChannel(conn) {
			fmt.Printf("%s\n", line)
		}

		fmt.Println("Connection closed")
	}
}
