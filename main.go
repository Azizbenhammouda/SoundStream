package main

import (
	"fmt"
	"log"
	"net"
	"slices"
	"strings"
)

func validateReq(req []string) bool {
	if len(req) != 3 {
		fmt.Println("Error: Invalid HTTP request line")
		return false
	}

	if !strings.HasPrefix(req[2], "HTTP/") {
		fmt.Printf("Error: Invalid HTTP protocol version '%s'\n", req[2])
		return false
	}

	methods := []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"HEAD",
		"OPTIONS",
		"PATCH",
	}

	if !slices.Contains(methods, req[0]) {
		fmt.Printf("Error: Invalid HTTP method '%s'\n", req[0])
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

	fmt.Println("GOX TCP server is running on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		buffer := make([]byte, 4096)

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Read error:", err)
			conn.Close()
			continue
		}

		if n == 0 {
			fmt.Println("Client sent an empty request")
			conn.Close()
			continue
		}

		request := string(buffer[:n])

		lines := strings.Split(request, "\r\n")

		if len(lines) == 0 {
			fmt.Println("Invalid HTTP request")
			conn.Close()
			continue
		}

		req := strings.Fields(lines[0])

		if !validateReq(req) {
			fmt.Println("Invalid HTTP request")
			conn.Close()
			continue
		}

		fmt.Printf(
			"method=%s path=%s version=%s\n",
			req[0],
			req[1],
			req[2],
		)

		body := "Hello from GOX"

		response := "HTTP/1.1 200 OK\r\n" +
			"Content-Length: 14\r\n" +
			"Content-Type: text/plain\r\n" +
			"\r\n" +
			body

		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Write error:", err)
		}

		conn.Close()
	}
}
