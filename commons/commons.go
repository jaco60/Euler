package commons

import "math"

func IsPrime(nb int) bool {
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
