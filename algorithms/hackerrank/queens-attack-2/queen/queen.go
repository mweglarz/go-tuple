package queen

import (
	"math"
)

type Pos struct {
	x int32
	y int32
}

func (self Pos) Equal(pos Pos) bool {
	return self.x == pos.x && self.y == pos.y
}

func (self Pos) isEmpty() bool {
	return self.x == 0 && self.y == 0
}

func (self Pos) Min() int32 {
	min := self.x
	if self.y < self.x {
		min = self.y
	}
	return min
}

func (self Pos) Max() int32 {
	max := self.x
	if self.y > self.x {
		max = self.y
	}
	return max
}

func QueenMaxDistance(queenPos Pos, obstacles [][]int32, size int32) int32 {
	obs := OptimizeObstacles(queenPos, obstacles, size)

	var dist int32 = 0

	// check colMin
	if queenPos.Equal(obs[0]) {
		dist += queenPos.y - 1

	} else {
		dist += distance(queenPos, obs[0])
	}

	// check colMax
	if queenPos.Equal(obs[1]) {
		dist += size - queenPos.y

	} else {
		dist += distance(queenPos, obs[1])
	}

	// check rowMin
	if queenPos.Equal(obs[2]) {
		dist += queenPos.x - 1

	} else {
		dist += distance(queenPos, obs[2])
	}

	// check rowMax
	if queenPos.Equal(obs[3]) {
		dist += size - queenPos.x

	} else {
		dist += distance(queenPos, obs[3])
	}

	// check SW
	if obs[4].isEmpty() {
		dist += queenPos.Min() - 1

	} else {
		dist += distance(queenPos, obs[4])
	}

	// check NW
	if obs[5].isEmpty() {
		dist += Pos{queenPos.x, size + 1 - queenPos.y}.Min() - 1

	} else {
		dist += distance(queenPos, obs[5])
	}

	// check NE
	if obs[6].isEmpty() {
		dist += Pos{size + 1 - queenPos.x, size + 1 - queenPos.y}.Min() - 1

	} else {
		dist += distance(queenPos, obs[6])
	}

	// check SE
	if obs[7].isEmpty() {
		dist += Pos{size + 1 - queenPos.x, queenPos.y}.Min() - 1

	} else {
		dist += distance(queenPos, obs[7])
	}

	return dist
}

func OptimizeObstacles(queenPos Pos, obstacles [][]int32, size int32) []Pos {

	var (
		queenColMin Pos = queenPos
		queenColMax Pos = queenPos
		queenRowMin Pos = queenPos
		queenRowMax Pos = queenPos
		queenDiagSW Pos = Pos{}
		queenDiagNW Pos = Pos{}
		queenDiagNE Pos = Pos{}
		queenDiagSE Pos = Pos{}
	)

	isDiagonal := isDiagonalGenerator(queenPos, size)
	isSecDiagonal := isSecDiagonalGenerator(queenPos, size)

	for _, posArr := range obstacles {

		pos := Pos{posArr[0], posArr[1]}
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
				if !queenDiagSW.isEmpty() {
					if pos.x > queenDiagSW.x {
						queenDiagSW = pos
					}
					continue

				} else {
					queenDiagSW = pos
					continue
				}
			}

			if pos.x > queenPos.x {
				if !queenDiagNE.isEmpty() {
					if pos.x < queenDiagNE.x {
						queenDiagNE = pos
					}
					continue

				} else {
					queenDiagNE = pos
					continue
				}
			}
		}

		if isSecDiagonal(pos) {
			if pos.x < queenPos.x {
				if !queenDiagNW.isEmpty() {
					if pos.x > queenDiagNW.x {
						queenDiagNW = pos
					}
					continue

				} else {
					queenDiagNW = pos
					continue
				}
			}

			if pos.x > queenPos.x {
				if !queenDiagSE.isEmpty() {
					if pos.x < queenDiagSE.x {
						queenDiagSE = pos
					}
					continue

				} else {
					queenDiagSE = pos
					continue
				}
			}
		}
	}

	return []Pos{
		queenColMin,
		queenColMax,
		queenRowMin,
		queenRowMax,
		queenDiagSW,
		queenDiagNW,
		queenDiagNE,
		queenDiagSE,
	}
}

func isDiagonalGenerator(queenPos Pos, size int32) func(Pos) bool {
	return func(pos Pos) bool {
		b := queenPos.y - queenPos.x
		return pos.y == (pos.x + b)
	}
}

func isSecDiagonalGenerator(queenPos Pos, size int32) func(Pos) bool {
	return func(pos Pos) bool {
		b := queenPos.y + queenPos.x
		return pos.y == (-pos.x + b)
	}
}

func distance(pos1, pos2 Pos) int32 {
	if pos1.x == pos2.x {
		abs := float64(pos1.y - pos2.y)
		return int32(math.Abs(abs)) - 1

	} else {
		abs := float64(pos1.x - pos2.x)
		return int32(math.Abs(abs)) - 1
	}
}
