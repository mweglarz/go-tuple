package bomberman

import (
	"crypto/md5"
	"encoding/json"
)

type Pos struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type Cell struct {
	Pos      `json:"pos"`
	Value    rune `json:"value"`
	BombTime int8 `json:"bombTime"`
}

func (self Cell) Hash() (string, error) {
	jsonBytes, err := json.Marshal(self)
	if err != nil {
		return "", err
	}
	bytes := md5.Sum(jsonBytes)
	return string(bytes[:]), nil
}

func BomberMan(n int32, grid []string) []string {

	return nil
}

func Transform(grid []string) [][]Cell {
	var cells [][]Cell
	var cellRow []Cell
	for row, line := range grid {
		cellRow = nil
		for col, ch := range line {
			pos := Pos{col, row}
			cellRow = append(cellRow, Cell{pos, ch})
		}
		cells = append(cells, cellRow)
	}
	return cells
}

func GetRealmInitState([][]Cell) [][]Cell {

}
