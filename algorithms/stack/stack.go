package stack

type Element interface{}

type EmptyStackError struct {
}

func (e EmptyStackError) Error() string {
	return "stack is empty"
}

type Stack struct {
	elements []Element
}

func New(initialValues ...Element) *Stack {
	return &Stack{initialValues}
}

func (st *Stack) Push(elem Element) {
	newStack := make([]Element, 0, len(st.elements)+1)
	newStack = append(newStack, elem)
	st.elements = append(newStack, st.elements...)
}

func (st *Stack) Pop() (result Element, err error) {
	if st.IsEmpty() {
		err = EmptyStackError{}
	}
	result = st.elements[0]
	st.elements = append(st.elements[:0], st.elements[1:]...)
	return
}

func (st *Stack) TopElement() (*Element, error) {
	if st.IsEmpty() {
		return nil, EmptyStackError{}
	}

	return &st.elements[0], nil
}

func (st *Stack) IsEmpty() bool {
	return len(st.elements) < 1
}
