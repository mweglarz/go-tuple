package mergeSort

type T interface{}

func Sorted(array []T, sorter func(T, T) bool) []T {
	if len(array) < 2 {
		return array
	}
	middleIndex := len(array) / 2

	leftPile := Sorted(array[:middleIndex], sorter)
	rightPile := Sorted(array[middleIndex:], sorter)

	return merge(leftPile, rightPile, sorter)
}

func merge(leftPile, rightPile []T, sorter func(T, T) bool) []T {
	leftCount := len(leftPile)
	rightCount := len(rightPile)
	orderedPile := make([]T, 0, leftCount+rightCount)

	leftIndex, rightIndex := 0, 0

	leftAction := func(value T) {
		orderedPile = append(orderedPile, value)
		leftIndex++
	}

	rightAction := func(value T) {
		orderedPile = append(orderedPile, value)
		rightIndex++
	}

	for leftIndex < leftCount && rightIndex < rightCount {
		var leftValue, rightValue T = leftPile[leftIndex], rightPile[rightIndex]

		if sorter(leftValue, rightValue) {
			leftAction(leftValue)

		} else if sorter(rightValue, leftValue) {
			rightAction(rightValue)

		} else {

			leftAction(leftValue)
			rightAction(rightValue)
		}
	}

	for leftIndex < leftCount {
		leftAction(leftPile[leftIndex])
	}

	for rightIndex < rightCount {
		rightAction(rightPile[rightIndex])
	}

	return orderedPile
}
