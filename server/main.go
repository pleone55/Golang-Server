package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	//setting a connection deadline
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Connection timeout")
	}
	//read from the connection
	scanner := bufio.NewScanner(conn)
	//returns a bool everytime it scans
	//fetches and prints each line of the host connection data
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		//write to a connection
		fmt.Fprintf(conn, "I heard: %s\n", ln)
	}
	defer conn.Close()

	//code never gets here initially
	//code gets her after setting a connection timeout
	fmt.Println("Now here")
}
