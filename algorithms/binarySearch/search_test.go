package binarySearch

import "testing"

// func TestSearch(t *testing.T) {
// 	numbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67}
// 	expect := 13
//
// 	index, _ := Search(numbers, 43)
//
// 	if expect != index {
// 		t.Fatalf("Expecting %d, found %d", expect, index)
// 	}
// }
//
// func TestSearchNearest(t *testing.T) {
// 	numbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67}
// 	expect := 6
//
// 	index, nearestIndex := Search(numbers, 18)
//
// 	if expect != nearestIndex {
// 		t.Fatalf("Expecting %d, found %d, nearest %d", expect, index, nearestIndex)
// 	}
// }

func TestEmpty(t *testing.T) {
	numbers := []int{}
	expect := -1

	index, _ := SearchDecreasing(numbers, 18)

	if expect != index {
		t.Fatalf("Expecting %d, found %d", expect, index)
	}
}

func TestAtEnd(t *testing.T) {
	numbers := []int{100, 50, 40}
	expect := 2

	index, _ := SearchDecreasing(numbers, 40)

	if expect != index {
		t.Fatalf("Expecting %d, found %d", expect, index)
	}
}

func Test1(t *testing.T) {
	numbers := []int{100, 50, 40, 20, 10}
	expect := 1

	index, _ := SearchDecreasing(numbers, 50)

	if expect != index {
		t.Fatalf("Expecting %d, found %d", expect, index)
	}

}
