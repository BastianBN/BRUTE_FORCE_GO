package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sync"
	"time"
	"unicode/utf8"
)

var possibilitescarac = "0123456789abcdefghijklmnopqrstuvwxyz"
var bytesdata, err = ioutil.ReadFile("mdp.txt")
var contenufichier = string(bytesdata)
var wg sync.WaitGroup
var retour = make(chan string)

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
	wg.Add(1)
	go recherche("", puissance)
	fmt.Println("MDP TROUVE : " + <-retour)
	wg.Wait()

	//for i, lettre := range possibilitescarac{
	//	wg.Add(1)
	//	fmt.Println("Goroutine numéro ", i)
	//	go recherche(string(lettre), int(possibilites)-1)
	//}
	//wg.Wait()

	tempspasse := time.Since(start)
	fmt.Println("Temps écoulé : ", tempspasse)
}

func recherche(chaine string, ncaracatrouver int) string {
	//fmt.Println(wg)
	if trysolution(chaine) {
		//fmt.Println("MDP TROUVE : " + chaine)
		wg.Done()
		retour <- chaine
		return chaine
	}
	for i := 0; i < 36; i++ {
		if ncaracatrouver == 0 {
			//fmt.Println(chaine)
			return chaine
		} else {
			//fmt.Println(i)
			//fmt.Println(ncaracatrouver)
			recherche(chaine+string(possibilitescarac[i]), ncaracatrouver-1)
		}
	}
	return ""
}

func trysolution(chaine string) bool {
	return chaine == contenufichier
}
