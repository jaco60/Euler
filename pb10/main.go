package main

import (
	"ProjetEuler/commons"
	"fmt"
)

//Quel est la somme de tous les nombres premiers en dessous de deux millions ?

func main() {
	const LIMITE = 2000000
	somme := 2
	for nb := 3; nb <= LIMITE; nb += 2 {
		if commons.IsPrime(nb) {
			somme += nb
		}
	}
	fmt.Println(somme)
}
