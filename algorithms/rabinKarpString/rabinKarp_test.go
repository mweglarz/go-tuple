package rabinKarp

import (
	"fmt"
	"testing"
)

func TestSearchPattern(t *testing.T) {
	fmt.Println("*** Starting TestSearchPattern ***")
	s := "The big dog jumped over the fox"
	p := "ump"
	expect := 13

	res, err := Search(s, p)
	if err != nil {
		t.Fatalf("Search returned error %v", err)
	}
	if res != expect {
		t.Fatalf("Returned index should be %d", expect)
	}
}

func TestCompareArrays(t *testing.T) {
	fmt.Println("*** Starting TestCompareArrays ***")
	ar1 := []int32{1, 2, 3, 4}
	ar2 := []int32{1, 2, 3, 4}

	if !compare(ar1, ar2) {
		t.Fatal("func compare doesn't work")
	}
}

func TestHash_And_NextHash(t *testing.T) {
	fmt.Println("*** Starting TestHash_And_NextHash ***")
	s1 := "aaaa"
	s1Array := ([]int32)(s1)
	hash1 := hash(s1Array[:3])
	nHash := nextHash(hash1, s1Array[0], s1Array[3], 2)

	if nHash != hash1 {
		t.Fatalf("hash1 = %f\ndifferent from\nnext hash =%f\n", hash1, nHash)
	}
}

func TestHash_And_NextHash2(t *testing.T) {
	fmt.Println("*** Starting TestHash_And_NextHash2 ***")
	s1 := "abcd"
	s1Array := ([]int32)(s1)
	hash1 := hash(s1Array[:3])
	nHash := nextHash(hash1, s1Array[0], s1Array[3], 2)
	hash2 := hash(s1Array[1:4])

	if nHash != hash2 {
		t.Fatalf("hash2 = %f\ndifferent from\nnext hash =%f\n", hash2, nHash)
	}
}
