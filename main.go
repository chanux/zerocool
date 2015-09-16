package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	port  string
	file  string
	delay int
)

func main() {
	flag.StringVar(&port, "p", "1337", "Port to listen on")
	flag.StringVar(&file, "f", "example.txt", "File to serve")
	flag.IntVar(&delay, "d", 50, "Delay between characters")
	flag.Parse()

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening: ", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error handling request: ", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	//msg := "It's working\n And yes"
	fmt.Println("A client starts reading")

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		char, err := reader.ReadByte()
		if err != nil {
			break
		}

		//conn.Write([]byte(string(each)))
		conn.Write([]byte{char})
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}

	fmt.Println("One client finishes reading")
	conn.Close()
}
