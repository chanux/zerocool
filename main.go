package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"strings"
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

	fmt.Println("connection from: ", conn.RemoteAddr())
	frtn, err := fortune(file)
	if err != nil {
		fmt.Println("failed to read fortune from file")
	}

	for i := 0; i < len(frtn); i++ {
		conn.Write([]byte{frtn[i]})
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}

	conn.Close()
}

// given a path representing a fortune file, load the file, parse it,
// an return a random fortune cookie
func fortune(fortuneFile string) (string, error) {
	content, err := ioutil.ReadFile(fortuneFile)
	var fortunes []string = nil
	if err == nil {
		fortunes = strings.Split(string(content), "%\n")
	}
	rand.Seed(time.Now().UTC().UnixNano())
	i := rand.Int() % len(fortunes)
	return fortunes[i], err
}
