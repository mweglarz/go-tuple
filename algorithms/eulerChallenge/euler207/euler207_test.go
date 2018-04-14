package euler207

import (
	"testing"
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
	_ = FindMulti([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10})
}

func checkLeastM(a, b, expect int, t *testing.T) {
	if m := FindLeastM(a, b); m != expect {
		t.Fatalf("incorrect least M, a = %d, b = %d, result = %d, expected = %d", a, b, m, expect)
	}
}
