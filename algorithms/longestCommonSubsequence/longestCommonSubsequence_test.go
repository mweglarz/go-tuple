package longestCommonSubsequence

import (
	"fmt"
	"testing"
)

func TestLCS(t *testing.T) {
	var s1, s2 = "abcbx", "abdcab"
	runTest(s1, s2, "abcb", t)
}

func TestLCS2(t *testing.T) {
	var s1, s2 = "abcbx", "klmk"
	runTest(s1, s2, "", t)
}

func TestLCS3(t *testing.T) {
	var s1, s2 = "Hello World", "Bonjour le monde"
	runTest(s1, s2, "oorld", t)
}

func runTest(s1, s2, expect string, t *testing.T) {
	if result, err := LCS(s1, s2); err != nil {
		t.Fatalf("Test failed, s1 = %s, s2 = %s", s1, s2)

	} else if result != expect {
		t.Fatalf("Expected %s", expect)

	} else {
		fmt.Println("Longest common subsequence:", result)
	}
}
