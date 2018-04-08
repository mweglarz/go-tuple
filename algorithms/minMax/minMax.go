package minMax

func MinMax(ar []int) (min, max int) {
	count := len(ar)
	if count == 0 {
		return
	} else if count == 1 {
		return ar[0], ar[0]
	}
	min, max = ar[0], ar[0]
	offset := count % 2
	for i := offset; i < count; i += 2 {
		minIndx, maxIndx := i, i+1
		if minFromPair(ar[i:i+2]) == i+1 {
			minIndx, maxIndx = i+1, i
		}
		if ar[minIndx] < min {
			min = ar[minIndx]
		}
		if ar[maxIndx] > max {
			max = ar[maxIndx]
		}
	}
	return
}

func minFromPair(pair []int) (index int) {
	index = 0
	if pair[1] < pair[0] {
		index = 1
	}
	return
}
