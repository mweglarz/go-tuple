package queen

type Pos struct {
	x int
	y int
}

func MaxQueenDistance(queenPos Pos, obstacles []Pos) int {

	var (
		queenColMin Pos = queenPos
		queenColMax Pos = queenPos
		queenRowMin Pos = queenPos
		queenRowMax Pos = queenPos
		queenDiag1  Pos = queenPos
		queenDiag2  Pos = queenPos
		queenDiag3  Pos = queenPos
		queenDiag4  Pos = queenPos
	)

	for _, pos := range obstacles {
		_ = queenColMin
		_ = queenColMax
		_ = queenRowMin
		_ = queenRowMax
		_ = queenDiag1
		_ = queenDiag2
		_ = queenDiag3
		_ = queenDiag4

		// same col
		if pos.x == queenPos.x {
			if pos.y > queenColMax.y {
				queenColMax = pos
			}
			if pos.y < queenColMin.y {
				queenColMin = pos
			}
		}
		// same row
		if pos.y == queenPos.y {
			if pos.x > queenRowMax.x {
				queenRowMax = pos
			}
			if pos.x < queenRowMin.x {
				queenRowMin = pos
			}
		}
	}

	return 0
}
