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

var mdp = "dabo"

var stop = 0
var possibilitescarac = "abcdefghijklmnopqrstuvwxyz0123456789"
var wglisteners sync.WaitGroup

func getArgs() int {
	res, _ := strconv.Atoi(os.Args[1])
	return res
}

func trysolution(chaine string) bool {
	return chaine == mdp
}

func main() {
	var portString string
	port := getArgs()
	portString = fmt.Sprintf("127.0.0.1:%d", port) //Modifier le 0 en i pour toutes les IP

	//    fmt.Printf("#DEBUG MAIN Creating TCP Server on port %d\n", ports[i])
	fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString)

	wglisteners.Add(1)
	ln, errListen := net.Listen("tcp", portString)
	if errListen != nil {
		fmt.Printf("ERROR on Listen : %s\n", errListen.Error())
	}
	go acceptConnection(ln)
	wglisteners.Wait()

}

func acceptConnection(ln net.Listener) {
	for {
		connum := 1
		fmt.Printf("#DEBUG MAIN Accepting next connection\n")
		conn, errconn := ln.Accept()

		if errconn != nil {
			fmt.Printf("DEBUG MAIN Error when accepting next connection\n")
			panic(errconn)
		}

		//If we're here, we did not panic and conn is a valid handler to the new connection

		fmt.Println("J'handle la co")
		go handleConnection(conn, connum)
		connum += 1
	}

}
func handleConnection(connection net.Conn, connum int) {
	defer connection.Close()
	connReader := bufio.NewReader(connection)
	inputLine, err := connReader.ReadString('\n')
	fmt.Println(inputLine)
	if err != nil {
		fmt.Printf("#DEBUG %d RCV ERROR no panic, just a client\n", connum)
		fmt.Printf("Error :|%s|\n", err.Error())
	} else {

		inputLine = strings.TrimSuffix(inputLine, "\n")
		puissance := utf8.RuneCountInString(mdp)
		fmt.Println(puissance)
		fmt.Println(inputLine)
		stop := 0
		retour := make(chan string)
		for i := 0; i < 36; i++ {
			fmt.Printf("Starting go routine with %s\n", inputLine+string(possibilitescarac[i]))
			go recherche(inputLine+string(possibilitescarac[i]), puissance-1, &stop, retour)

		}

		a := ""
		for i := 0; i < 36; i++ {
			a = <-retour
			fmt.Printf("Got return [%s] on channel\n", a)
			if len(a) != 0 {
				_, _ = io.WriteString(connection, fmt.Sprintf("MDP_FOUND %s\n", a))
				break
			}
		}
		if len(a) == 0 {
			_, _ = io.WriteString(connection, fmt.Sprintf("MDP_NOTFOUND\n", a))
		}
		fmt.Printf("All jobs done, go routines over")
	}
}
func recherche(chaine string, ncaracatrouver int, stop *int, retour chan string) {
	retour <- rechercheAux(chaine, ncaracatrouver, stop, retour)
}

func rechercheAux(chaine string, ncaracatrouver int, stop *int, retour chan string) string {
	//fmt.Println(chaine)
	//fmt.Println(ncaracatrouver)
	if *stop == 1 {
		return ""
	} else {

		if trysolution(chaine) {
			fmt.Println("MDP TROUVE : " + chaine)
			fmt.Println(runtime.NumGoroutine())
			*stop = 1
			return chaine
		} else {
			if ncaracatrouver == -1 {
				return ""
			} else {
				for i := 0; i < 36; i++ {
					res := rechercheAux(chaine+string(possibilitescarac[i]), ncaracatrouver-1, stop, retour)
					if res != "" {
						return res
					}
				}
				return ""
			}
		}
	}
}
