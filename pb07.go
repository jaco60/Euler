package main

import (
	"fmt"
)

// Quel est le 10001ème nombre premier ?

func main() {
	cpt, nb := 1, 1
	for cpt < 10001 {
		nb += 2
		if isPrime(nb) {
			cpt++
		}
	}
	fmt.Println(nb)
}
