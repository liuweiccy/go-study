package test

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Server() {
	fmt.Println("Starting the server ...")
	listener, err := net.Listen("tcp", "localhost:50010")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}
func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Println("Received data:", string(buf))
	}
}

func Client() {
	conn, err := net.Dial("tcp", "localhost:50010")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("what's your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")
	for {
		fmt.Println("what to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
		if err != nil {
			return
		}
	}
}
