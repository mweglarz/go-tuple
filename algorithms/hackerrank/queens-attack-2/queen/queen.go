package queen

type Pos struct {
	x int
	y int
}

func MaxQueenDistance(queenPos Pos, obstacles []Pos) int {

	var (
		queenColMin Pos  = queenPos
		queenColMax Pos  = queenPos
		queenRowMin Pos  = queenPos
		queenRowMax Pos  = queenPos
		queenDiagSW *Pos = nil
		queenDiagNW *Pos = nil
		queenDiagNE *Pos = nil
		queenDiagSE *Pos = nil
	)

	isDiagonal := func(pos Pos) bool {
		return false
	}
	isSecDiagonal := func(pos Pos) bool {
		return false
	}

	for _, pos := range obstacles {

		// same col
		if pos.x == queenPos.x {
			if pos.y > queenColMax.y {
				queenColMax = pos
				continue
			}
			if pos.y < queenColMin.y {
				queenColMin = pos
				continue
			}
		}
		// same row
		if pos.y == queenPos.y {
			if pos.x > queenRowMax.x {
				queenRowMax = pos
				continue
			}
			if pos.x < queenRowMin.x {
				queenRowMin = pos
				continue
			}
		}

		if isDiagonal(pos) {
			if pos.x < queenPos.x {
				if queenDiagSW != nil {
					if pos.x > queenDiagSW.x {
						queenDiagSW = &pos
					}
					continue

				} else {
					queenDiagSW = &pos
					continue
				}
			}

			if pos.x > queenPos.x {
				if queenDiagNE != nil {
					if pos.x < queenDiagNE.x {
						queenDiagNE = &pos
					}
					continue

				} else {
					queenDiagNE = &pos
					continue
				}
			}
		}

		if isSecDiagonal(pos) {
			if pos.x < queenPos.x {
				if queenDiagNW != nil {
					if pos.x > queenDiagNW.x {
						queenDiagNW = &pos
					}
					continue

				} else {
					queenDiagNW = &pos
					continue
				}
			}

			if pos.x > queenPos.x {
				if queenDiagSE != nil {
					if pos.x < queenDiagSE.x {
						queenDiagSE = &pos
					}
					continue

				} else {
					queenDiagSE = &pos
					continue
				}
			}
		}
	}

	return 0
}

func distance(pos1, pos2 Pos) int {
	return 0
}
