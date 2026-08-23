package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("TCP server is on")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		buffer := make([]byte, 1024)
		n, err2 := conn.Read(buffer) //n is the number of bytes actually received
		if err2 != nil {
			fmt.Println(err2.Error())
			conn.Close()
			continue
		}
		fmt.Println(string(buffer[:n]))

		conn.Close()

	}
}
