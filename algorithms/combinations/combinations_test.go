package combinations

import (
	"testing"
)

func TestCombinations(t *testing.T) {

	array := []int32{1, 2, 3, 4, 5}
	k := 2
	expected := [][]int32{
		{1, 2}, {1, 3}, {1, 4}, {1, 5},
		{2, 3}, {2, 4}, {2, 5},
		{3, 4}, {3, 5},
		{4, 5}}

	result := Combinations(array, k)

	for i := 0; i < len(expected); i++ {
		if expected[i][0] != result[i][0] || expected[i][1] != result[i][1] {
			t.Fatalf("Result doesn't match expected, result = %v,\nexpected = %v", result, expected)
		}
	}
}
