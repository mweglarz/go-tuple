package repeatingVariations

import (
	"fmt"
	"testing"
)

func TestVariations(t *testing.T) {
	fmt.Println("****** Starting test", "TestVariations", "******")

	pattern := "hej"
	FindVariations(4, pattern, "")
}
