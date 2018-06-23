package subset

func NonDivisibleSubset(k int32, array []int32) int32 {

	count := int32(0)
	counts := make([]int32, k)
	for _, v := range array {
		key := v % k
		counts[key] = counts[key] + 1
	}

	if counts[0] > 0 {
		count++
	}

	for i := int32(1); i < (k/2)+1; i++ {
		j := k - i
		if i != j {
			toAdd := counts[i]
			if counts[j] > counts[i] {
				toAdd = counts[j]
			}
			count += toAdd
		}
	}
	if k%2 == 0 {
		count++
	}

	return count
}
