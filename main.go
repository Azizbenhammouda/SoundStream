package main

import (
	"fmt"
	"log"
	"net"
	"slices"
	"strings"
)

func validateReq(req []string) bool {
	if len(req) < 3 {
		fmt.Println("Error: Invalid HTTP request line length")
		return false
	}
	if !strings.HasPrefix(req[2], "HTTP/") {
		fmt.Printf("Error: Invalid HTTP protocol version '%s'\n", req[2])
		return false
	}
	methods := []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS", "PATCH"}
	if !slices.Contains(methods, req[0]) {
		fmt.Println("Error: Invalid HTTP method")
		return false
	}

	return true
}

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
		n, err := conn.Read(buffer) // n is the number of bytes actually received
		if err != nil {
			fmt.Println(err.Error())
			conn.Close()
			continue
		}
		if n == 0 {
			fmt.Println("client sent an empty message")
			conn.Close()
			continue
		}
		request := string(buffer[:n])

		lines := strings.Split(request, "\r\n")
		re := strings.Fields(lines[0])

		if len(lines) == 0 || !validateReq(re) {
			fmt.Println("Invalid HTTP request")
			conn.Close()
			continue
		}

		req := strings.Fields(lines[0])

		fmt.Printf("method=%s path=%s version=%s\n", req[0], req[1], req[2])
		conn.Close()
	}
}
