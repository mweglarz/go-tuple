package bomberman

import (
	"fmt"
	"testing"
)

var testGrid []string = []string{
	"..O",
	"...",
	"...",
}

func TestTransform(t *testing.T) {

	result := Transform(testGrid)
	fmt.Println(result)
	fmt.Println(AsciiView(result))
	hash := Hash(result)
	fmt.Println(hash)
}

func TestHashing(t *testing.T) {
	c1 := Cell{Pos{9, 1}, 'd', EMPTY}
	c2 := Cell{Pos{9, 2}, 'd', EMPTY}

	h1, _ := c1.Hash()
	h2, _ := c2.Hash()
	fmt.Println(h1)
	fmt.Println(h2)
}
