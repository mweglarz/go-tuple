package reversePolishNotation

import (
	"fmt"
	"testing"
)

type BasicOperation struct {
	char      string
	priority  int
	operation func(float64, float64) float64
}

func (op BasicOperation) GetPriority() int {
	return op.priority
}

func (op BasicOperation) IsDoubleArgument() bool {
	return true
}

func (op BasicOperation) Execute(args ...float64) float64 {
	return op.operation(args[0], args[1])
}

func (op BasicOperation) String() string {
	return op.char
}

func TestChangeToNotationTest(t *testing.T) {
	fmt.Println("****** Starting test", "TestChangeToNotationTest", "******")

	var plus Operation = BasicOperation{"+", 100, func(a, b float64) float64 { return a + b }}
	var prod Operation = BasicOperation{"*", 200, func(a, b float64) float64 { return a * b }}

	operations := []interface{}{float64(1), plus, float64(2), plus, float64(3), prod, float64(4), plus, float64(3), prod, float64(6), prod, float64(7)}

	result := ConvertToPolishNotation(operations)
	fmt.Printf("operations = %+v\n", operations)
	fmt.Printf("result = %+v\n", result)
}
