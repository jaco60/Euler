package main

import (
	"fmt"
	"math"
)

//Quel est la somme de tous les nombres premiers en dessous de deux millions ?

func isPrime(nb int) bool {
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	root := int(math.Sqrt(float64(nb)))
	for div := 3; div <= root; div = div + 2 {
		if nb%div == 0 {
			return false
		}
	}
	return true
}

func main() {
	const LIMITE = 2000000
	somme := 2
	for nb := 3; nb <= LIMITE; nb += 2 {
		if isPrime(nb) {
			somme += nb
		}
	}
	fmt.Println(somme)
}
