package main

import "fmt"

// Somme de tous les multiples de 3 ou 5 inférieurs à 1000
func main() {
	somme := 0
	for n := 3; n < 1000; n++ {
		if n%3 == 0 || n%5 == 0 {
			somme += n
		}
	}
	fmt.Println(somme)
}
