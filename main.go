package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	connection, err := net.Dial("tcp", "localhost:8888")
	logFatal(err)

	defer connection.Close()

	log.Printf("connected!")

	for {
		go read(connection)
		write(connection)
	}
}

func read(connection net.Conn) {
	reader := bufio.NewReader(connection)
	message, err := reader.ReadString('\n')

	if err == io.EOF {
		connection.Close()
		fmt.Println("Connection closed.")
		os.Exit(0)
	}

	fmt.Println(message)
}

func write(connection net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	connection.Write([]byte(message))
}
