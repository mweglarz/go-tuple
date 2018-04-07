package euler207

import (
	"fmt"
	"math"
)

func FindLeastM(a, b int) int {

	lim := float64(a) / float64(b)

	var k int = 1
	for t := 1; frac(t, k) >= lim; t++ {
		var presentPerfectK int = maxKForNaturalT(t)
		var nextPerfectK int = maxKForNaturalT(t + 1)
		k = nextPerfectK - 1
		fmt.Printf("k = %+v\n", k)

		if frac(t, k) < lim {
			// solved, now find least
			for i := k - 1; i >= presentPerfectK; i-- {
				fmt.Printf("i = %+v\n", i)

				// TODO: go backwars and check frac, maybe do binary search here??
				if frac(t, i) >= lim {
					// this frac is too big, return previous one
					return i + 1
				}
			}
		}
	}
	return -1
}

func frac(num, den int) float64 {
	result := float64(num) / float64(den)
	fmt.Printf("num = %d, den = %d, result = %f\n", num, den, result)
	return result
}

func maxKForNaturalT(t int) int {
	return int(math.Pow(4, float64(t))) - int(math.Pow(2, float64(t)))
}

func computeT(k int) float64 {
	d := math.Sqrt(1 + 4*float64(k))
	t := math.Log2(1+d) - 1
	return t
}

func isTNatural(t float64) bool {
	tN := int(t)
	if float64(tN)-t < 10e-6 {
		return true
	}
	return false
}
