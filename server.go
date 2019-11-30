package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func getArgs() [4]int {
	var portNumbers [4]int
	var err error
	null := [4]int{-1, -1, -1, -1}
	if len(os.Args) != 5 {
		fmt.Printf("Usage: go run client.go <4 portnumbers>\n")
		os.Exit(1)
	} else {
		fmt.Printf("#DEBUG ARGS Port Numbers : %s\n", os.Args[1:5])
		for i := 1; i <= 4; i++ {
			//fmt.Println(os.Args[i])
			portNumbers[i-1], err = strconv.Atoi(os.Args[i])
			if err != nil {
				fmt.Printf("Usage: go run client.go <4 portnumbers>\n")
				os.Exit(1)
			}
		}
		return portNumbers
	}
	//Should never be reached

	return null
}

func main() {
	var portString string
	ports := getArgs()
	for i := 0; i < 4; i++ {
		fmt.Printf("#DEBUG MAIN Creating TCP Server on port %d\n", ports[i])
		portString = fmt.Sprintf(":%s", strconv.Itoa(ports[i]))
		fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString)
	}

	ln, err := net.Listen("tcp", portString)
	if err != nil {
		fmt.Printf("#DEBUG MAIN Could not create listener\n")
		panic(err)
	}

	//If we're here, we did not panic and ln is a valid listener

	connum := 1

	for {
		fmt.Printf("#DEBUG MAIN Accepting next connection\n")
		conn, errconn := ln.Accept()

		if errconn != nil {
			fmt.Printf("DEBUG MAIN Error when accepting next connection\n")
			panic(errconn)
		}

		//If we're here, we did not panic and conn is a valid handler to the new connection

		go handleConnection(conn, connum)
		connum += 1

	}
}

func handleConnection(connection net.Conn, connum int) {

	defer connection.Close()
	connReader := bufio.NewReader(connection)
	//    if err !=nil{
	//        fmt.Printf("#DEBUG %d handleConnection could not create reader\n", connum)
	//        return
	//    }

	for {
		inputLine, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Printf("#DEBUG %d RCV ERROR no panic, just a client\n", connum)
			fmt.Printf("Error :|%s|\n", err.Error())
			break
		}

		//fmt.Printf("#DEBUG RCV |%s|\n", inputLine)
		inputLine = strings.TrimSuffix(inputLine, "\n")
		fmt.Printf("#DEBUG %d RCV |%s|\n", connum, inputLine)
	}

}
