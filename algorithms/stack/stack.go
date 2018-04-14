package stack

import (
	"errors"
)

type Element interface{}

type stack struct {
	elements []Element
}

func New(initialValues ...Element) *stack {
	return &stack{initialValues}
}

func (st *stack) Push(elem Element) {
	newStack := make([]Element, 0, len(st.elements)+1)
	newStack = append(newStack, elem)
	st.elements = append(newStack, st.elements...)
}

func (st *stack) Pop() (result Element, err error) {
	if len(st.elements) < 1 {
		err = errors.New("stack is empty")
		return
	}
	result = st.elements[0]
	st.elements = append(st.elements[:0], st.elements[1:]...)
	return
}
