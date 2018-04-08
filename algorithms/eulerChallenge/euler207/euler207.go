package euler207

import (
	"fmt"
	"math"
)

// FIXME: 4^t and 2^t have to be integers !!!
func FindLeastM(a, b int) int {

	lim := float64(a) / float64(b)

	var k int = 1
	for perfectT := 1; frac(perfectT, k) >= lim; perfectT++ {
		var presentPerfectK int = maxKForNaturalT(perfectT)
		var nextPerfectK int = maxKForNaturalT(perfectT + 1)
		k = nextPerfectK - 1
		fmt.Printf("k = %+v\n", k)
		t := computeT(k)

		findLeastFromSolved := func(solvedK int) int {
			// solved, now find least
			for i := k - 1; i >= presentPerfectK; i-- {
				if _, err := computTwo(t); err != nil {
					continue
				}

				// TODO: go backwars and check frac, maybe do binary search here??
				if frac(perfectT, i) >= lim {
					// this frac is too big, return previous one
					return i + 1
				}
			}
			return solvedK
		}

		for i := k; i >= presentPerfectK; i-- {
			t = computeT(i)
			if _, err := computTwo(t); err == nil && frac(perfectT, i) < lim {
				return findLeastFromSolved(i)
			}
		}
	}
	return -1
}

func P(m int) float64 {
	var all float64 = 0
	var perfect float64 = 0
	for k := 1; k <= m; k++ {
		fmt.Printf("P(%d) checking solution for k = %d\n", m, k)
		t := computeT(k)
		if _, err := computTwo(t); err == nil {
			all += 1
			fmt.Printf("P(%d) found solution for k = %d\n", m, k)
			if _, ok := isInteger(t); ok {
				perfect += 1
				fmt.Printf("P(%d) solution for k = %d\n, t = %f is perfect", m, k, t)
			}
		}
	}
	return perfect / all
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

func computTwo(t float64) (int, error) {
	two := math.Pow(2, t)
	fmt.Println("two =", two)
	if intTwo, ok := isInteger(two); ok {
		return intTwo, nil
	}
	return 0, fmt.Errorf("not integer")
}

func isInteger(num float64) (int, bool) {
	n := int(num + 0.5)
	if compare(num, float64(n)) {
		return n, true
	}
	return 0, false
}

func compare(a, b float64) bool {
	return math.Abs(a-b) < 10e-6
}
