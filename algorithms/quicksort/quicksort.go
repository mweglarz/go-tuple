package quicksort

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
