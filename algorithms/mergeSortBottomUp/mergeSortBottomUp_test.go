package mergeSortBottomUp

import (
	"fmt"
	"testing"
)

func TestSort_ArrayOfInt_Returns_SortedArray(t *testing.T) {
	inputArray := []T{2, 1, 5, 4, 9}
	sortedArray := []int{1, 2, 4, 5, 9}

	inputArray = Sorted(inputArray, func(l, r T) bool {
		fmt.Println("l value", l, "right value", r)
		return l.(int) < r.(int)
	})

	if len(inputArray) != len(sortedArray) {
		t.Error("Arrays length mismatch")
	}

	for i := 0; i < len(inputArray); i++ {
		if inputArray[i].(int) != sortedArray[i] {
			t.Errorf("value %d from input array: %v \n different from %v \n",
				inputArray[i].(int), inputArray, sortedArray)
		}
	}
}
