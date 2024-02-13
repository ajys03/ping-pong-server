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

}

func main() {
	fmt.Println("Starting Ping Pong Application...")

	ln, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		fmt.Print(err)
	}
	// TODO: might have to defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Print(err)
		}
		go handleConnection(conn)
	}
}
