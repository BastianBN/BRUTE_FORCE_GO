package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
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

func lancementServers(portString [36]string) [36]net.Conn {
	var connections [36]net.Conn
	var err error
	possibilitescarac := "abcdefghijklmnopqrstuvwxyz0123456789"
	for i := 0; i < 36; i++ {
		fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString[i])
		connections[i], err = net.Dial("tcp", portString[i])
		if err != nil {
			fmt.Printf("#DEBUG MAIN could not connect\n")
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println(string(possibilitescarac[i]))
			_, _ = io.WriteString(connections[i], string(possibilitescarac[i])+"\n")
		}
	}
	return connections
}

func lecture(connection net.Conn, wg sync.WaitGroup) {
	reader := bufio.NewReader(connection)

	resultString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("DEBUG MAIN could not read from server")
		os.Exit(1)
	}
	resultString = strings.TrimSuffix(resultString, "\n")
	fmt.Printf("#DEBUG server replied : |%s|\n", resultString)
	wg.Done()
}

func main() {
	ports := getArgs()
	fmt.Println(ports)
	var listeIp [9]string //Définition des adresses IP auxquelles on va demander une connexion.
	listeIp[0] = "127.0.0.1"
	//listeIp[1] =
	//listeIp[2] =
	//listeIp[3] =
	//listeIp[4] =
	//listeIp[5] =
	//listeIp[6] =
	//listeIp[7] =
	//listeIp[8] =

	var portString [36]string //On prépare les strings IP:PORT à l'avance.
	for i := 0; i < 9; i++ {
		for j := 0; j < 4; j++ {
			portString[i*4+j] = fmt.Sprintf("%s:%s", listeIp[i], strconv.Itoa(ports[j])) //Modifier le 0 en i pour toutes les IP
		}
	}
	fmt.Println(portString)
	var wg sync.WaitGroup
	wg.Add(36)
	fmt.Printf("#DEBUG MAIN connected\n")
	connections := lancementServers(portString)

	for i := 0; i < 36; i++ {
		go lecture(connections[i], wg)
	}

	time.Sleep(1000 * time.Millisecond)
}
