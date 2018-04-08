package minMax

import (
	"fmt"
	"testing"
)

func TestMinMax(t *testing.T) {
	fmt.Println("****** Starting test", "TestMinMax", "******")
	ar := []int{2, 3, 1, 5, 6, 3, 9}

	min, max := MinMax(ar)
	if min != 1 && max != 9 {
		t.Fatalf("test failed, min = %d, max = %d", min, max)
	}
}
