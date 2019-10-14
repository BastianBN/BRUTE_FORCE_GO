package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
)

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

	caracteres := "0123456789abcdefghijklmnopqrstuvwxyz"

	fmt.Println(caracteres)
	fmt.Println(reflect.TypeOf(caracteres))
	for i := 0; i < 36; i++ {
		fmt.Println(string(caracteres[i]))
	}

	bytesdata, err := ioutil.ReadFile("mdp.txt")
	fichier := string(bytesdata)
	if err != nil {
		fmt.Println(err)
	}

	oui := "jod"
	fmt.Println(fichier)
	fmt.Println(oui)
	if oui == fichier {
		fmt.Println("nice")
	} else {
		fmt.Println("woops")
	}
	//fmt.Scanln()
}
