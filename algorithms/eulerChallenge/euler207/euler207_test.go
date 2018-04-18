package euler207

import (
	"testing"

	"tuple-mw.com/algorithms/eulerChallenge/euler207/find"
)

// func TestLeastM_1(t *testing.T) {
// 	fmt.Println("****** Starting test", "TestLeastM1", "******")
//
// 	checkLeastM(2, 3, 6, t)
// 	checkLeastM(9, 20, 30, t)
// 	checkLeastM(1, 10, 2550, t)
// 	checkLeastM(1, 20, 14520, t)
// }

func TestLeastMultiM(t *testing.T) {
	// _ = FindMulti([]int{2, 9, 1, 1}, []int{3, 20, 10, 20})
	//_ = FindMulti([]int{1, 1}, []int{10, 10})
	_ = FindMulti([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10})
	// _ = FindMulti([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2})
}

func TestFindNearestLowerSolution(t *testing.T) {
	ar := []int{1, 6, 12, 25}
	k := 7

	indx, _ := find.FindNearestLowerSolution(k, ar, 0, len(ar))
	if ar[indx] != 6 {
		t.Fatalf("incorrect nearest solution, ar[indx] = %d", ar[indx])
	}
}

func TestFindNearestLargerSolution_SameValue(t *testing.T) {
	ar := []int{1, 6, 12, 25}
	k := 6

	indx, _ := find.FindNearestLowerSolution(k, ar, 0, len(ar))
	if ar[indx] != 6 {
		t.Fatalf("incorrect nearest solution, ar[indx] = %d", ar[indx])
	}
}

func checkLeastM(a, b, expect int, t *testing.T) {
	if m := FindLeastM(a, b); m != expect {
		t.Fatalf("incorrect least M, a = %d, b = %d, result = %d, expected = %d", a, b, m, expect)
	}
}
