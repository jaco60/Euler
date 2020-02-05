package main

import (
	"ProjetEuler/commons"
	"fmt"
	"math"
	"sort"
)

// Quel est le plus grand facteur premier du nombre 600851475143 ?

func factors(nbre int) []int {
	res := []int{}
	root := int(math.Sqrt(float64(nbre)))
	for d := 2; d < root; d++ {
		if nbre%d == 0 {
			res = append(append(res, d), nbre/d)
		}
	}
	sort.Ints(res)
	if nbre == root*root {
		return res[:len(res)-1]
	} else {
		return res
	}
}

func main() {
	facts := factors(600851475143)
	i := len(facts) - 1
	for i >= 0 && !commons.IsPrime(facts[i]) {
		i--
	}
	fmt.Printf("Plus grand facteur premier de %d : %d\n", 600851475143, facts[i])
}
