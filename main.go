package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
	flag.StringVar(&file, "f", "fortunes.txt", "File to serve")
	flag.IntVar(&delay, "d", 50, "Delay between characters")
	flag.Parse()

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("error handling request: %v", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	log.Println("connection from: ", conn.RemoteAddr())
	frtn, err := fortune(file)
	if err != nil {
		log.Printf("failed to read fortune: %v\n", err)
	}

	for i := 0; i < len(frtn); i++ {
		_, err := conn.Write([]byte{frtn[i]})
		if err != nil {
			log.Print(err)
			break
		}
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}

	conn.Close()
}

// given a path representing a fortune file, load the file, parse it,
// an return a random fortune cookie
func fortune(fortuneFile string) (string, error) {
	content, err := ioutil.ReadFile(fortuneFile)
	if err != nil {
		return "", fmt.Errorf("failed to read from file: %v", err)
	}
	fortunes := strings.Split(string(content), "%\n")
	rand.Seed(time.Now().UTC().UnixNano())
	i := rand.Int() % len(fortunes)
	return fortunes[i], nil
}
