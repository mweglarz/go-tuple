package encrypt

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	testString := "haveaniceday"
	expectedResult := "hae and via ecy"

	result := Encrypt(testString)

	if result != expectedResult {
		t.Fatal("results don't match")
	}

}
