package euler207

import (
	"fmt"
	"math"
)

func FindLeastM(a, b int) int {
	var result float64 = 1.0
	i := 2
	for ; result > frac(a, b); i++ {
		result = P(i)
	}
	return i - 1
}

func P(m int) float64 {
	var all float64 = 0
	var perfect float64 = 0
	for k := 1; k <= m; k++ {
		t := computeT(k)
		if _, err := computTwo(t); err == nil {
			all += 1
			if _, ok := isInteger(t); ok {
				perfect += 1
			}
		}
	}
	return perfect / all
}

func frac(num, den int) float64 {
	result := float64(num) / float64(den)
	return result
}

func computeT(k int) float64 {
	d := math.Sqrt(1 + 4*float64(k))
	t := math.Log2(1+d) - 1
	return t
}

func computTwo(t float64) (int, error) {
	two := math.Pow(2, t)
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
