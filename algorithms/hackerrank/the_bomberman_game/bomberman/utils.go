package bomberman

import "fmt"

func Transform(grid []string) [][]Cell {
	var cells [][]Cell
	var cellRow []Cell
	for _, line := range grid {
		cellRow = nil
		for _, ch := range line {
			cellRow = append(cellRow, Cell{ch, GetState(ch)})
		}
		cells = append(cells, cellRow)
	}
	return cells
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
		rowArray := make([]rune, len(cells[i]))
		for j := 0; j < len(cells[i]); j++ {
			cell := cells[i][j]
			rowArray[j] = cell.GetChar()
		}
		rows = append(rows, string(rowArray))
	}
	return rows
}

// debug only
func PrintState(cells [][]Cell) {
	ascii := AsciiView(cells)
	for _, line := range ascii {
		fmt.Println(line)
	}
}
