package bomberman

import (
	"crypto/md5"
	"encoding/json"
)

type BombState uint8

const (
	PLANTED BombState = iota
	WAIT1
	WAIT2
	EXPLODING
	EMPTY
)

var time int = 0
var hashes map[string]int
var period int32

type Pos struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Cell struct {
	Pos   `json:"pos"`
	Value rune      `json:"value"`
	State BombState `json:"bombTime"`
}

func (self *Cell) TimeProgress() {
	if self.State != EMPTY {
		self.State += 1
	}
}

func (self *Cell) PlantBombIfEmpty() {
	if self.State == EMPTY {
		self.State = PLANTED
	}
}

func (self *Cell) Hash() (string, error) {
	jsonBytes, err := json.Marshal(self)
	if err != nil {
		return "", err
	}
	bytes := md5.Sum(jsonBytes)
	return string(bytes[:]), nil
}

func Hash(cells [][]Cell) string {
	result := ""

	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			cellHash, _ := cells[i][j].Hash()
			result = result + cellHash
		}
	}
	return result
}

func Detonate(cells [][]Cell) {
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			cell := cells[i][j]
			// TODO: to implement
		}
	}
}

func GetState(ch rune) BombState {
	if ch == 79 {
		return PLANTED
	}
	return EMPTY
}

func BomberMan(n int32, grid []string) []string {
	cells := Transform(grid)
	AdvanceToInitState(cells)
	hashExist := func() {
		_, ok := hashes[Hash(cells)]
		return ok
	}
	for !hashExist() {
		time++
		if time%2 == 1 {
			Detonate(cells)

		} else {
			ForEach(cells, func(c *Cell) {
				c.PlantBombIfEmpty()
			})
		}
		hashes[Hash(cells)] = time
	}
	period = time - 3

	return nil
}

func Transform(grid []string) [][]Cell {
	var cells [][]Cell
	var cellRow []Cell
	for row, line := range grid {
		cellRow = nil
		for col, ch := range line {
			pos := Pos{col, row}
			cellRow = append(cellRow, Cell{pos, ch, GetState(ch)})
		}
		cells = append(cells, cellRow)
	}
	return cells
}

func AdvanceToInitState(cells [][]Cell) {

	ForEach(cells, func(c *Cell) {
		c.TimeProgress()
		c.TimeProgress()
		c.PlantBombIfEmpty()
		c.TimeProgress()
	})
	time += 3
	Detonate(cells)
	hashes[Hash(cells)] = time

	return cells
}

func ForEach(cells [][]Cell, action func(cell *Cell)) {
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			action(&cells[i][j])
		}
	}
}
