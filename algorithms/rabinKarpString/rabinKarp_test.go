package rabinKarp

import "testing"

func TestSearchPattern(t *testing.T) {
	s := "The big dog jumped over the fox"
	p := "ump"
	expect := 13

	res, err := Search(s, p)
	if err != nil {
		t.Fatal("Search returned error")
	}
	if res != expect {
		t.Fatalf("Returned index should be %d", expect)
	}
}
