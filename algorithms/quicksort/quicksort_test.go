package quicksort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	fmt.Println("****** Starting test", "TestSort", "******")

	testArray := []int{10, 0, 3, 9, 2, 14, 8, 27, 1, 5, 8, -1, 26}
	mockSorted := []int{-1, 0, 1, 2, 3, 5, 8, 8, 9, 10, 14, 26, 27}
	sorted := Sorted(testArray)
	DutchSort(testArray, 0, len(testArray)-1)
	fmt.Printf("testArray = %+v\n", testArray)

	for i := 0; i < len(testArray); i++ {
		if sorted[i] != mockSorted[i] {
			t.Fatalf("Wrong sorting at index %d", i)
		}
	}
	fmt.Println(Sorted(testArray))
}
