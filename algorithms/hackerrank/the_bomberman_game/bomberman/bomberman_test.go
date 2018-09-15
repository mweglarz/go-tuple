package bomberman

import (
	"fmt"
	"testing"
)

var testGrid []string = []string{

	".......",
	"...O...",
	"....O..",
	".......",
	"OO.....",
	"OO.....",
}

var initTestGrid []string = []string{

	"OOO.OOO",
	"OO...OO",
	"OOO...O",
	"..OO.OO",
	"...OOOO",
	"...OOOO",
}

func TestTransform(t *testing.T) {

	result := Transform(testGrid)
	// fmt.Println(result)
	asciiResult := AsciiView(result)
	// fmt.Println(asciiResult)

	if equal, msg := compare(asciiResult, testGrid); !equal {
		t.Errorf(msg)
	}
	// hash := Hash(result)
	// fmt.Println(hash)
}

func compare(grid []string, grid2 []string) (bool, string) {

	if len(grid) != len(grid2) {
		return false, "Length don't match"
	}

	for i := 0; i < len(grid); i++ {
		if grid[i] != grid2[i] {
			msg := fmt.Sprintf("%v not equal %v", []rune(grid[i]), []rune(grid2[i]))
			return false, msg
		}
	}
	return true, ""
}
