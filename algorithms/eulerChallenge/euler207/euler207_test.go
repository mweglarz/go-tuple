package euler207

import (
	"fmt"
	"math"
	"testing"
)

func TestLeastM_1(t *testing.T) {
	fmt.Println("****** Starting test", "TestLeastM1", "******")

	for i := 1; i < 5; i++ {
		k := maxKForNaturalT(i)
		fmt.Printf("max k = %d, for t = %d\n", k, i)
	}
	checkLeastM(2, 3, 6, t)
	checkLeastM(9, 20, 30, t)
}

func TestFirst3K(t *testing.T) {
	fmt.Println("****** Starting test", "TestFirst3K", "******")
	if !(checkBackwards(1) && checkBackwards(2) && checkBackwards(3)) {
		t.Fatal("checking backwards failed")
	}
	checkBackwards(4)
	checkBackwards(5)
	checkBackwards(6)
}

func TestMaxKForSession(t *testing.T) {
	fmt.Println("****** Starting test", "TestMaxKForSession", "******")
	maxK := maxKForNaturalT(1)
	if maxK != 2 {
		t.Fatal("maxK should be 2")
	}

	maxK = maxKForNaturalT(2)
	if maxK != 12 {
		t.Fatal("maxK should be 12")
	}
}

func checkLeastM(a, b, expect int, t *testing.T) {
	if m := FindLeastM(a, b); m != expect {
		t.Fatalf("incorrect least M, a = %d, b = %d, result = %d, expected = %d", a, b, m, expect)
	}
}

func checkBackwards(k int) bool {
	t := computeT(k)
	fmt.Printf("t = %+v\n", t)
	four := math.Pow(4, t)
	fmt.Printf("four = %+v\n", four)

	computedK := math.Pow(4, t) - math.Pow(2, t)
	diff := computedK - float64(k)

	if diff < 10e-6 {
		return true
	}
	return false
}
