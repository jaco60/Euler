package main

// Quel est le plus petit entier positif pouvant être divisé par chacun des nombres entre 1 et 20 sans aucun reste ?

func main() {

	for nb := 20; ; nb++ {
		ok := true
		for div := 1; div <= 20; div++ {
			if nb%div != 0 {
				ok = false
				break
			}
		}
		if ok {
			print(nb)
			break
		}
	}
}
