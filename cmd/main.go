package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to bind to port 6379")
		os.Exit(1)
	}
	fmt.Println("Started Accepting Conenctions......")
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Fprint(os.Stderr, "Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(c)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
			return
		}
		if bytes.Equal([]byte("*1\\r\\n$4\\r\\nPING\\r\\n"), buffer[:n-1]) {
			pong := []byte("+PONG\\r\\n")
			conn.Write(pong)
		}
	}

}
