package stack

import (
	"fmt"
	"testing"
)

func TestStackBasicFunctionality(t *testing.T) {
	fmt.Println("****** Starting test", "TestStackBasicFunctionality", "******")
	stack := New()
	stack.Push(4)
	fmt.Printf("stack.elements = %+v\n", stack.elements)
	stack.Push(3)
	fmt.Printf("stack.elements = %+v\n", stack.elements)
	popped, _ := stack.Pop()
	fmt.Printf("stack.elements = %+v\npopped = %d\n", stack.elements, popped)

	if popped.(int) != 3 {
		t.Fatalf("expecting 3 as popped result get %d", popped.(int))
	}
}
