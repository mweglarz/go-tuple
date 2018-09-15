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

// func TestHashing(t *testing.T) {
// c1 := Cell{Pos{9, 1}, 'd', EMPTY}
// c2 := Cell{Pos{9, 2}, 'd', EMPTY}

// h1, _ := c1.Hash()
// h2, _ := c2.Hash()
// fmt.Println(h1)
// fmt.Println(h2)
// }

func TestAdvanceToInitState(t *testing.T) {

	initCells := Transform(testGrid)

	AdvanceToInitState(initCells)
	initGrid := AsciiView(initCells)

	if equal, msg := compare(initGrid, initTestGrid); !equal {
		t.Errorf(msg)
	}
}

func TestDetonated(t *testing.T) {


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
