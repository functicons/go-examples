package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

// TCP server which serves a finite or infinite text stream of numbers.
//
// Usage: go run text_stream.go <port> <sendInternalMs> <maxCount>

func handleConnection(connection net.Conn, interval int, maxCount int) {
	var remoteAddress = connection.RemoteAddr().String()

	defer func() {
		fmt.Printf("Closing connection with %s\n", remoteAddress)
		connection.Close()
	}()

	fmt.Printf("Handling connection from %s\n", remoteAddress)
	for i := 0; maxCount == -1 || i < maxCount; i = i + 1 {
		var line = strconv.Itoa(i) + "\n"
		fmt.Printf("Sending: %v", line)
		var _, err = connection.Write([]byte(line))
		if err != nil {
			fmt.Printf("Error sending: %v", line)
			return
		}
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func main() {
	var port int = 8080
	var sendIntervalMs int = 1000
	var maxCount = -1
	var err error

	// Parse command-line args.
	if len(os.Args) >= 2 {
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Invalid port number: %v\n", os.Args[1])
			os.Exit(1)
		}
	}
	if len(os.Args) >= 3 {
		sendIntervalMs, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Invalid send interval (ms): %v\n", os.Args[2])
			os.Exit(2)
		}
	}
	if len(os.Args) >= 4 {
		maxCount, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("Invalid max count: %v\n", os.Args[3])
			os.Exit(3)
		}
	}
	fmt.Printf(
		"Arguments: port=%d, sendIntervalMs=%d, maxCount=%d\n",
		port,
		sendIntervalMs,
		maxCount)

	// Listen.
	var address = "0.0.0.0:" + strconv.Itoa(port)
	fmt.Printf("Listening on %v\n", address)
	listener, err := net.Listen("tcp4", address)
	if err != nil {
		fmt.Printf("Error listening: %v\n", err)
		os.Exit(4)
	}

	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting: %v\n", err)
			continue
		}
		go handleConnection(connection, sendIntervalMs, maxCount)
	}
}
