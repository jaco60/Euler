package main

import (
	"fmt"
)

// Quel est le plus petit entier positif pouvant être divisé par chacun des nombres entre 1 et 20 sans aucun reste ?

// Par force brute...
func force_brute() int {
	for nb := 20; ; nb++ {
		ok := true
		for div := 1; div <= 20; div++ {
			if nb%div != 0 {
				ok = false
				break
			}
		}
		if ok {
			return nb
		}
	}
}

/* On renvoie le plus petit commun multiple des nbres de 1 à 20
Les facteurs premiers des nombres de 1 à 20 sont :
1 : 1, 2 : 2, 3 : 3, 4 : 2^2, 5 : 5, 6 : 2 * 3, 7 : 7, 8 : 2^3, 9 : 3^2, 10 : 2 * 5, 11 : 11, 12 : 2^2 * 3,
13 : 13, 14 : 2 * 7, 15 : 3 * 5, 16 : 2^4, 17 : 17, 18 : 2 * 3^2, 19 : 19, 20 : 2^2 * 5

Les facteurs premiers sont donc 2^4, 3^2, 5, 7, 11, 13, 17 et 19...
*/

func ppcm() int {
	return 2 * 2 * 2 * 2 * 3 * 3 * 5 * 7 * 11 * 13 * 17 * 19
}

func main() {
	fmt.Println(force_brute())
	fmt.Println(ppcm())
}
