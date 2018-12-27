package main

import "testing"

func TestSample2(t *testing.T) {

	s1 := "SHINCHAN"
	s2 := "NOHARAAA"
	expect := int32(3)

	test(s1, s2, expect, t)
}

func Test1(t *testing.T) {

	s1 := "WEWOUCUIDGCGTRMEZEPXZFEJWISRSBBSYXAYDFEJJDLEBVHHKS"
	s2 := "FDAGCXGKCTKWNECHMRXZWMLRYUCOCZHJRRJBOAJOQJZZVUYXIC"
	expect := int32(15)

	test(s1, s2, expect, t)
}

func test(s1, s2 string, expect int32, t *testing.T) {

	result := commonChild(s1, s2)

	if result != expect {
		t.Fatalf("Result was %d, expected %d", result, expect)
	}
}
