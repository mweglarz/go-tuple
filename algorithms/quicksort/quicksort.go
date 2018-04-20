package quicksort

import (
	"math/rand"
	"time"
)

type sorter func(int) bool

func Sorted(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	mid := len(ar) / 2
	pivot := ar[mid]
	less := filter(ar, func(a int) bool {
		return a < pivot
	})
	equal := filter(ar, func(a int) bool {
		return a == pivot
	})
	greater := filter(ar, func(a int) bool {
		return a > pivot
	})

	less = Sorted(less)
	greater = Sorted(greater)

	return append(less, append(equal, greater...)...)
}

func filter(ar []int, fun sorter) []int {
	var result []int
	for _, v := range ar {
		if fun(v) {
			result = append(result, v)
		}
	}
	return result
}

func DutchSort(ar []int, low, hi int) {
	if low < hi {
		pivotIndx := randInt(low, hi)
		p, q := partitionDutchFlag(ar, low, hi, pivotIndx)
		DutchSort(ar, low, p-1)
		DutchSort(ar, q+1, hi)
	}
}

func partitionDutchFlag(ar []int, low, hi, pivotIndx int) (int, int) {
	var (
		pivot   = ar[pivotIndx]
		smaller = low
		equal   = low
		larger  = hi
	)

	for equal <= larger {
		if ar[equal] < pivot {
			swap(ar, smaller, equal)
			smaller++
			equal++

		} else if ar[equal] == pivot {
			equal++

		} else {
			swap(ar, equal, larger)
			larger--
		}
	}
	return smaller, larger
}

func swap(ar []int, aIndx, bIndx int) {
	tmp := ar[aIndx]
	ar[aIndx] = ar[bIndx]
	ar[bIndx] = tmp
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
