package encrypt

import (
	"testing"
)

func TestEncryptSample1(t *testing.T) {
	testString := "haveaniceday"
	expectedResult := "hae and via ecy"

	test(testString, expectedResult, t)
}

func TestEncrypt3(t *testing.T) {
	testString := "iffactsdontfittotheorychangethefacts"
	expectedResult := "isieae fdtonf fotrga anoyec cttctt tfhhhs"

	test(testString, expectedResult, t)
}

func test(testString, expectedResult string, t *testing.T) {

	result := Encrypt(testString)

	if result != expectedResult {
		t.Fatalf("results don't match, expected = %s, got = %s",
			expectedResult, result)
	}
}
