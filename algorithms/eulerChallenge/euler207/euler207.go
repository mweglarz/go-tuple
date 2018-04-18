package euler207

import (
	"fmt"
	"math"

	"tuple-mw.com/algorithms/eulerChallenge/euler207/find"
)

var results = make(map[int]float64)
var kWithResult = make([]int, 0)

func FindMulti(as []int, bs []int) []int {
	var result []int
	for i := 0; i < len(as); i++ {
		result = append(result, FindLeastM(as[i], bs[i]))
	}
	return result
}

func FindLeastM(a, b int) int {
	var tresh float64 = frac(a, b)
	var result float64 = 1.0
	var i int = 2

	for ; result > tresh; i++ {
		result = P(i)
	}
	return i - 1
}

func P(m int) float64 {

	if len(kWithResult) > 0 && kWithResult[len(kWithResult)-1] >= m {

		indx, _ := find.FindNearestLowerSolution(m, kWithResult, 0, len(kWithResult))
		k := kWithResult[indx]
		value := results[k]
		return value
	}

	var all float64 = 0
	var perfect float64 = 0
	for k := 1; k <= m; k++ {
		t := computeT(k)
		if _, err := computTwo(t); err == nil {
			appendToKsIfNeeded(k)
			all += 1
			results[k] = perfect / all

			if _, ok := isInteger(t); ok {
				perfect += 1
				results[k] = perfect / all
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

func appendToKsIfNeeded(k int) {
	_, err := find.Find(k, kWithResult, 0, len(kWithResult))
	if err != nil {
		kWithResult = append(kWithResult, k)
	}
}
