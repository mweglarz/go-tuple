package binarySearch

// Return index of element -1 if not found,
// second returned value is last midIndex or -1
func SearchDecreasing(array []int, value int) (int, int) {

	var lowerBound, upperBound = 0, len(array)
	var midIndex int = -1
	for lowerBound < upperBound {
		midIndex = lowerBound + (upperBound-lowerBound)/2

		if array[midIndex] == value {
			return midIndex, -1

		} else if array[midIndex] > value {
			lowerBound = midIndex + 1

		} else {
			upperBound = midIndex
		}
	}

	return -1, midIndex
}
