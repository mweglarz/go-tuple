package reducedString

import (
	"testing"
)

func TestReduceSample1(t *testing.T) {
	str := "aaabccddd"
	expect := "abd"
	out := Reduce(str)
	if out != expect {
		t.Errorf("expected %s, got %s", expect, out)
	}
}
