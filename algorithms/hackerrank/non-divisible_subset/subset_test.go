package subset

import "testing"

func TestSample(t *testing.T) {

	k := int32(3)
	array := []int32{1, 7, 2, 4}

	expect := int32(3)
	if result := NonDivisibleSubset(k, array); result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}
