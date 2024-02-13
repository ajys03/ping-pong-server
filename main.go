package main

import "fmt"
import "net"

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Print(err)
		}
	}(conn)

	buf := make([]byte, 4)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err.Error())
		return
	}
	if string(buf[:4]) != "ping" {
		fmt.Println("Received unexpected data:", string(buf))
		return
	}

	response := []byte("pong\n")
	_, err = conn.Write(response)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	return
}

func main() {
	fmt.Println("Starting Ping Pong Application...")

	ln, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Listening on Port 1080...")
	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ln)
	for {
		// ln.accept blocks till new connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println("Accepted connection from:", conn.RemoteAddr().String())
		go handleConnection(conn)
	}
}
