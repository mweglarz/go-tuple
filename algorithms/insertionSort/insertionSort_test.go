package insertionSort

import (
	"testing"
)

func TestSort_ArrayOfInt_Returns_SortedArray(t *testing.T) {
	inputArray := []T{8, 3, 2, 1, 5, 3, 9, 7}
	var sortedArray = []int{1, 2, 3, 3, 5, 7, 8, 9}
	Sort(inputArray, func(l, r T) bool {
		return l.(int) < r.(int)
	})

	if len(inputArray) != len(sortedArray) {
		t.Error("Arrays length mismatch")
	}

	for i := 0; i < len(inputArray); i++ {
		if inputArray[i].(int) != sortedArray[i] {
			t.Errorf("input array: %v \n different from %v \n", inputArray, sortedArray)
		}
	}
}

func TestSort_ArrayOfString_Returns_SortedArray(t *testing.T) {

	inputArray := []T{"dupa", "test", "ala", "ola"}
	var sortedArray = []string{"ala", "dupa", "ola", "test"}
	Sort(inputArray, func(l, r T) bool {
		return l.(string) < r.(string)
	})

	if len(inputArray) != len(sortedArray) {
		t.Error("Arrays length mismatch")
	}
	for i := 0; i < len(inputArray); i++ {
		if inputArray[i].(string) != sortedArray[i] {
			t.Errorf("input array: %v \n different from %v \n", inputArray, sortedArray)
		}
	}
}
