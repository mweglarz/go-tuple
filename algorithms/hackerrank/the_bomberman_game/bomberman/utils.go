package bomberman

import "fmt"

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

func PrintState(cells [][]Cell) {
	ascii := AsciiView(cells)
	for _, line := range ascii {
		fmt.Println(line)
	}
}
