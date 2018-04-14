package euler207

import (
	"fmt"
	"math"
)

type Result struct {
	k     int
	value float64
}

var lastValidK int = 0
var lastValidK2 int = 0
var switcher bool = false
var perfectResults = make([]Result, 0)

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
	for indx, res := range perfectResults {
		fmt.Printf("res = %+v\n", res)

		if res.value < tresh {
			i = perfectResults[indx-1].k
		}
	}
	for ; result > tresh; i++ {
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
			switcher = !switcher
			if switcher {
				lastValidK = k
			} else {
				lastValidK2 = k
			}
			if _, ok := isInteger(t); ok {
				var result Result
				if switcher {
					result = Result{k: lastValidK2, value: (perfect / (all - 1))}

				} else {
					result = Result{k: lastValidK, value: (perfect / (all - 1))}
				}
				perfectResults = append(perfectResults, result)
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
