package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

var mdp = "allo"
var retour = make(chan string)
var stop = 0
var possibilitescarac = "abcdefghijklmnopqrstuvwxyz0123456789"
var wg sync.WaitGroup
var waitingGroupListeners sync.WaitGroup

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

func trysolution(chaine string) bool {
	return chaine == mdp
}

func main() {
	var portString [36]string
	ports := getArgs()
	for i := 0; i < 4; i++ {
		portString[i] = fmt.Sprintf(":%s", strconv.Itoa(ports[i])) //Modifier le 0 en i pour toutes les IP
	}

	for i := 0; i < 4; i++ {
		fmt.Printf("#DEBUG MAIN Creating TCP Server on port %d\n", ports[i])
		fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString[i])
	}

	var listener [4]net.Listener
	for i := 0; i < 4; i++ {
		listener[i] = listenConnection(portString, i)
	}

	waitingGroupListeners.Add(4)
	go acceptConnection(listener[0])
	go acceptConnection(listener[1])
	go acceptConnection(listener[2])
	go acceptConnection(listener[3])
	wg.Wait()
}

func listenConnection(portString [36]string, indexDuPort int) net.Listener {
	ln, err := net.Listen("tcp", portString[indexDuPort])
	if err != nil {
		fmt.Printf("#DEBUG MAIN Could not create listener\n")
		panic(err)
	}
	return ln
}

func acceptConnection(ln net.Listener) {
	connum := 1
	fmt.Printf("#DEBUG MAIN Accepting next connection\n")
	conn, errconn := ln.Accept()

	if errconn != nil {
		fmt.Printf("DEBUG MAIN Error when accepting next connection\n")
		panic(errconn)
	}

	//If we're here, we did not panic and conn is a valid handler to the new connection

	fmt.Println("J'handle la co")
	fmt.Println(connum)
	go handleConnection(conn, connum)
	connum += 1

}
func handleConnection(connection net.Conn, connum int) {
	defer connection.Close()
	fmt.Println("oui salut on est dans le handle")
	connReader := bufio.NewReader(connection)
	inputLine, err := connReader.ReadString('\n')
	for {
		fmt.Println(inputLine)
		if err != nil {
			fmt.Printf("#DEBUG %d RCV ERROR no panic, just a client\n", connum)
			fmt.Printf("Error :|%s|\n", err.Error())
			break
		}

		inputLine = strings.TrimSuffix(inputLine, "\n")
		puissance := utf8.RuneCountInString(mdp)
		fmt.Println(puissance)

		wg.Add(1)
		fmt.Println(inputLine)
		go recherche(inputLine, puissance-1)

		_, _ = io.WriteString(connection, fmt.Sprintf("MDP TROUVE : %s \n", <-retour))
		break
	}
}

func recherche(chaine string, ncaracatrouver int) string {
	//fmt.Println(chaine)
	//fmt.Println(ncaracatrouver)
	if trysolution(chaine) {
		//fmt.Println("MDP TROUVE : " + chaine)
		fmt.Println(runtime.NumGoroutine())
		//for i := 0; i < runtime.NumGoroutine(); i++{
		//	fmt.Println(wg)
		//	wg.Done()
		//}
		fmt.Println("STOP")
		stop = 1
		retour <- chaine
		close(retour)
		//return chaine
	}
	for i := 0; i < 36; i++ {
		if stop == 0 {
			if ncaracatrouver == 0 {
				//fmt.Println(chaine)
				return chaine
			} else {
				//fmt.Println(i)
				//fmt.Println(chaine + string(possibilitescarac[i]))
				recherche(chaine+string(possibilitescarac[i]), ncaracatrouver-1)
			}
		}
	}
	return ""
}
