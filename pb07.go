package main

import (
	"fmt"
	"math"
)

// sub-optimal...
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
