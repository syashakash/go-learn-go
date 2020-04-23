package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

/*
	To test run nc localhost 8080 and start typing press ENTER to view response
 */

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("tcp server listener error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("tcp server accept error:", err)
		}

		go handleConnection(conn)
	}
}


func handleConnection(conn net.Conn) {
	bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		log.Println("client left")
		conn.Close()  // imp to close connection
		return
	}

	message := string(bufferBytes)
	clientAddr := conn.RemoteAddr().String()
	response := fmt.Sprintf(message + " from " + clientAddr + "\n")

	log.Println(response)

	conn.Write([]byte("you sent: " + response))

	handleConnection(conn)  // to ensure that client left gets printed if connection lost
}