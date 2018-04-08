package tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	fmt.Println("****** Starting test", "TestTree", "******")

	tree := New(3)
	a := New(5)
	a2 := New(6)
	a3 := New(2)
	a4 := New(9)
	a5 := New(8)

	tree.AddChild(a)
	tree.AddChild(a2)
	a2.AddChild(a3)
	a2.AddChild(a4)
	a4.AddChild(a5)

	if found := tree.Search(9); found != nil {
		fmt.Printf("found value!")
	} else {
		t.Fatal("not found")
	}
}
