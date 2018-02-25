package main

import "fmt"

// Il n’existe qu’un seul triplet pythagoricien pour lequel a + b + c = 1000. Déterminez le produit abc de celui-ci.
// Triplet pythagoricien (a, b, c) tel que a^2 + b^2 = c^2 avec a < b < c

func main() {
	const SOMME = 1000
	var a, b, c int
loop:
	for a = 1; a <= SOMME/3; a++ {
		for b = a; b <= SOMME/2; b++ {
			c = SOMME - a - b
			if a*a+b*b == c*c {
				break loop
			}
		}
	}
	fmt.Println(a, b, c)

}
