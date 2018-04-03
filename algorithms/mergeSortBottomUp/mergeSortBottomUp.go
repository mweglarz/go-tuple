package mergeSortBottomUp

type T interface{}

func Sorted(array []T, sorter func(T, T) bool) []T {
	if len(array) < 2 {
		return array
	}
	count := len(array)
	buffArray := make([]T, count, count)
	copy(buffArray, array)
	buff := [][]T{array, buffArray}
	d := 0

	for width := 1; width < count; width, d = 2*width, 1-d {
		for i := 0; i < count; i += 2 * width {
			var j, leftIndex, rightIndex int = i, i, i + width

			leftMax := min(leftIndex+width, count)
			rightMax := min(rightIndex+width, count)

			for ; leftIndex < leftMax && rightIndex < rightMax; j++ {
				if sorter(buff[d][leftIndex], buff[d][rightIndex]) {
					buff[1-d][j] = buff[d][leftIndex]
					leftIndex += 1

				} else {
					buff[1-d][j] = buff[d][rightIndex]
					rightIndex += 1
				}
			}

			for ; leftIndex < leftMax; j, leftIndex = j+1, leftIndex+1 {
				buff[1-d][j] = buff[d][leftIndex]
			}
			for ; rightIndex < rightMax; j, rightIndex = j+1, rightIndex+1 {
				buff[1-d][j] = buff[d][rightIndex]
			}
		}
	}

	return buff[d]
}

func min(l, r int) int {
	result := l
	if r < l {
		result = r
	}
	return result
}
