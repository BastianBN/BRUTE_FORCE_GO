package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"runtime"
	"sync"
	"time"
	"unicode/utf8"
)

var possibilitescarac = "abcdefghijklmnopqrstuvwxyz0123456789"
var bytesdata, err = ioutil.ReadFile("mdp.txt")
var contenufichier = string(bytesdata)
var wg sync.WaitGroup
var retour = make(chan string)
var stop = 0

func main() {
	/*fmt.Print("Nombre de caractères dans le mdp à cracker : ")
	var puissance int
	_, err := fmt.Scanf("%d", &puissance, 64)
	if err != nil {
		print(err)
	}*/
	// bytesdata, err := ioutil.ReadFile("mdp.txt")
	// fichier := string(bytesdata)
	if err != nil {
		fmt.Println(err)
	}
	puissance := utf8.RuneCountInString(contenufichier)
	fmt.Println(puissance)

	possibilites := math.Pow(36, float64(puissance))
	fmt.Println(possibilites)

	start := time.Now()

	//wg.Add(1)
	//go recherche("", puissance)
	//fmt.Println("MDP TROUVE : " + <-retour)
	//wg.Wait()

	for i := 0; i < 4; i++ {
		fmt.Println(possibilitescarac[i*9 : 9+i*9])
	}
	for i := 0; i < 36; i++ {
		wg.Add(1)
		carac := string(possibilitescarac[i])
		fmt.Println(carac)
		fmt.Println("Goroutine numéro ", i)
		go recherche(carac, puissance-1)
	}
	fmt.Println("MDP TROUVE : " + <-retour)
	tempspasse := time.Since(start)
	fmt.Println("Temps écoulé : ", tempspasse)
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

func trysolution(chaine string) bool {
	return chaine == contenufichier
}
