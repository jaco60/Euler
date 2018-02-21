package main

import (
	"fmt"
	"strconv"
)

// Quel est le plus grand palindrome créé à partir du produit de deux nombres à 3 chiffres ?

func isPalin(nbre int) bool {
	digits := strconv.Itoa(nbre)
	i, j := 0, len(digits)-1
	for i < j && digits[i] == digits[j] {
		i++
		j--
	}
	return i >= j
}

func main() {
	max := 0
	for x := 100; x <= 999; x++ {
		for y := x; y <= 999; y++ {
			prod := x * y
			if isPalin(prod) && prod > max {
				max = prod
			}
		}
	}
	fmt.Println(max)

}
