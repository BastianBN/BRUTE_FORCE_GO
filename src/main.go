package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"time"
)

var possibilitescarac = "0123456789abcdefghijklmnopqrstuvwxyz"

func main() {
	//scanner := bufio.NewReader(os.Stdin) //On utilise le paquet bufio pour créer un scanner qui va lire l'entrée dans la console.
	fmt.Print("Nombre de caractères dans le mdp à cracker : ")
	var puissance int
	_, err := fmt.Scanf("%d", &puissance, 64)
	//input, _ := scanner.ReadString('\n') // \n = utilisateur appuyant sur entrée, on pourrait mettre " " pour espace.
	if err != nil {
		print(err)
	}

	fmt.Println(puissance)
	fmt.Println(reflect.TypeOf(puissance))

	possibilites := math.Pow(36, float64(puissance))
	fmt.Println(possibilites)
	fmt.Println(possibilitescarac)
	fmt.Println(reflect.TypeOf(possibilitescarac))

	//fmt.Println(string(caracteres[i]))

	bytesdata, err := ioutil.ReadFile("mdp.txt")
	fichier := string(bytesdata)
	if err != nil {
		fmt.Println(err)
	}
	start := time.Now()
	recherche("", puissance, fichier)
	tempspasse := time.Since(start)
	fmt.Println("Temps écoulé : ", tempspasse)

	//fmt.Println(fichier)

	/*var compteur int
	for i := 0; i < 36; i++ {
		for j := 0; j < 36; j++{
			for k := 0; k < 36; k++{
				for l := 0; l < 36; l++{
					for m := 0; m < 36; m++ {
						chaineatester := string(caracteres[i]) + string(caracteres[j]) + string(caracteres[k]) + string(caracteres[l]) + string(caracteres[m])
						fmt.Println(chaineatester)
						compteur += 1
						if chaineatester == fichier {
							return
						}
					}
				}
			}
		}
	}
	fmt.Println(compteur)*/
	//fmt.Scanln()

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
