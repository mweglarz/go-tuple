package reversePolishNotation

import (
	"tuple-mw.com/algorithms/stack"
)

type Operation interface {
	GetPriority() int
	IsDoubleArgument() bool
	Execute(args ...float64) float64
	String() string
}

type operationList []interface{}

func ConvertToPolishNotation(operations operationList) operationList {
	result := make([]interface{}, 0)
	stack := stack.New()

	for _, v := range operations {
		switch v.(type) {
		case float64:
			result = append(result, v.(float64))
		case Operation:
			result = handleOperation(v.(Operation), stack, result)
		default:
			panic("Unknown type while processing OperationList")
		}
	}
	return result
}

func handleOperation(operation Operation, st *stack.Stack, result operationList) operationList {
	topEl, _ := st.TopElement()
	if topEl != nil && (*topEl).(Operation).GetPriority() >= operation.GetPriority() {
		popped, _ := st.Pop()
		result = append(result, popped)
		return handleOperation(operation, st, result)

	} else {
		st.Push(operation)
	}

	return result
}
