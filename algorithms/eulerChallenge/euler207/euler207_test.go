package euler207

import (
	"fmt"
	"testing"
)

func TestLeastM_1(t *testing.T) {
	fmt.Println("****** Starting test", "TestLeastM1", "******")

	checkLeastM(2, 3, 6, t)
	checkLeastM(9, 20, 30, t)
}

func checkLeastM(a, b, expect int, t *testing.T) {
	if m := FindLeastM(a, b); m != expect {
		t.Fatalf("incorrect least M, a = %d, b = %d, result = %d, expected = %d", a, b, m, expect)
	}
}
