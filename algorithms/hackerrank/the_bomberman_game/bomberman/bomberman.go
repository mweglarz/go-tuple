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
var period int

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

func (self *Cell) GetChar() rune {
	if self.State == EMPTY {
		return '.'
	}
	return 79
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

			if cells[i][j].State == EXPLODING {

				if i != 0 && cells[i-1][j].State != EXPLODING {
					cells[i-1][j].State = EMPTY
				}

				if i != len(cells)-1 && cells[i+1][j].State != EXPLODING {
					cells[i+1][j].State = EMPTY
				}

				if j != 0 && cells[i][j-1].State != EXPLODING {
					cells[i][j-1].State = EMPTY
				}

				if j != len(cells[i])-1 && cells[i][j+1].State != EXPLODING {
					cells[i][j+1].State = EMPTY
				}
				cells[i][j].State = EMPTY
			}
		}
	}
}

func GetState(ch rune) BombState {
	if ch == 79 {
		return PLANTED
	}
	return EMPTY
}

func AsciiView(cells [][]Cell) []string {
	var rows []string
	for i := 0; i < len(cells); i++ {
		rowArray := make([]rune, len(cells))
		for j := 0; j < len(cells[i]); j++ {
			cell := cells[i][j]
			rowArray = append(rowArray, cell.GetChar())
		}
		rows = append(rows, string(rowArray))
	}
	return rows
}

func BomberMan(n int32, grid []string) []string {
	cells := Transform(grid)
	AdvanceToInitState(cells)
	hashExist := func() bool {
		_, ok := hashes[Hash(cells)]
		return ok
	}
	simulate := func() {
		time++
		if time%2 == 1 {
			Detonate(cells)

		} else {
			ForEach(cells, func(c *Cell) {
				c.PlantBombIfEmpty()
			})
		}

	}
	for !hashExist() {
		hashes[Hash(cells)] = time
		simulate()
	}
	period = time - 3
	simulatesTodo := int(n) % period
	for i := 0; i < simulatesTodo; i++ {
		simulate()
	}

	return AsciiView(cells)
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
}

func ForEach(cells [][]Cell, action func(cell *Cell)) {
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			action(&cells[i][j])
		}
	}
}
