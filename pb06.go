package main

// Quelle est la différence entre la somme des carrés des cent premiers entiers naturels et le carré de leur somme ?

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func main() {
	somme_carres, somme := 0, 0
	for n := 1; n <= 100; n++ {
		somme_carres += n * n
		somme += n
	}
	println(abs(somme_carres - somme*somme))
}
