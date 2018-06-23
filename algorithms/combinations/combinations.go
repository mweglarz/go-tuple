package combinations

func Combinations(array []int32, k int) [][]int32 {

	var result [][]int32
	n := len(array)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			var pair []int32 = make([]int32, 2)
			pair[0], pair[1] = array[i], array[j]
			result = append(result, pair)
		}
	}

	return result
}
