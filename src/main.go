package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"time"
	"unicode/utf8"
)

var possibilitescarac = "0123456789abcdefghijklmnopqrstuvwxyz"

func main() {
	/*fmt.Print("Nombre de caractères dans le mdp à cracker : ")
	var puissance int
	_, err := fmt.Scanf("%d", &puissance, 64)
	if err != nil {
		print(err)
	}*/
	bytesdata, err := ioutil.ReadFile("mdp.txt")
	fichier := string(bytesdata)
	if err != nil {
		fmt.Println(err)
	}
	puissance := utf8.RuneCountInString(fichier)

	fmt.Println(puissance)
	fmt.Println(reflect.TypeOf(puissance))

	possibilites := math.Pow(36, float64(puissance))
	fmt.Println(possibilites)

	start := time.Now()
	recherche("", puissance, fichier)
	tempspasse := time.Since(start)
	fmt.Println("Temps écoulé : ", tempspasse)
}

func recherche(chaine string, ncaracatrouver int, contenufichier string) string {

	if chaine == contenufichier {
		fmt.Println("MDP TROUVE : " + chaine)
	}
	for i := 0; i < 36; i++ {
		if ncaracatrouver == 0 {
			//fmt.Println(chaine)
			return chaine
		} else {
			//fmt.Println(i)
			//fmt.Println(ncaracatrouver)
			recherche(chaine+string(possibilitescarac[i]), ncaracatrouver-1, contenufichier)
		}
	}
	return ""
}
