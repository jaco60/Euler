package main

import "fmt"

// Somme des termes pairs de Fibonacci < 4000000
func main() {
	a, b, sum := 1, 1, 0

	for b <= 4_000_000 {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}

	fmt.Println(sum)

}
