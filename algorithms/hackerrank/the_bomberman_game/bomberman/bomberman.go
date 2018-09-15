package bomberman

var time int32 = 0
var hashes map[string]int32 = make(map[string]int32)
var period int32

func BomberMan(n int32, grid []string) []string {
	cells := Transform(grid)

	hashLength := len(cells) * len(cells[0])

	hashExist := func(hash string) bool {
		_, ok := hashes[hash]
		return ok
	}

	hash := make([]rune, hashLength)

	simulate := func() {
		hash = make([]rune, hashLength)
		time++

		for i := 0; i < len(cells); i++ {
			for j := 0; j < len(cells[i]); j++ {

				cells[i][j].TimeProgress()

				// detonate
				if time%2 == 1 && cells[i][j].State == EXPLODING {

					if i != 0 && cells[i-1][j].State != WAIT2 {
						cells[i-1][j].UpdateState(EMPTY)
					}

					if i != len(cells)-1 && cells[i+1][j].State != EXPLODING {
						cells[i+1][j].UpdateState(EMPTY)
					}

					if j != 0 && cells[i][j-1].State != WAIT2 {
						cells[i][j-1].UpdateState(EMPTY)
					}

					if j != len(cells[i])-1 && cells[i][j+1].State != EXPLODING {
						cells[i][j+1].UpdateState(EMPTY)
					}
					cells[i][j].UpdateState(EMPTY)

					// plant
				} else if time%2 == 0 {
					cells[i][j].PlantBombIfEmpty()
				}

				hashIndex := i*len(cells[i]) + j
				hash[hashIndex] = cells[i][j].GetChar()
			}
		}

		// after detonating
		if time%2 == 1 {
			hashes[string(hash)] = time
		}
	}

	for !hashExist(string(hash)) {
		if time == 0 {
			time++
			for i := 0; i < len(cells); i++ {
				for j := 0; j < len(cells[i]); j++ {
					cells[i][j].TimeProgress()
				}
			}

		} else if time == 1 {
			time++
			for i := 0; i < len(cells); i++ {
				for j := 0; j < len(cells[i]); j++ {
					cells[i][j].TimeProgress()
					cells[i][j].PlantBombIfEmpty()
				}
			}

		} else {
			simulate()
		}

		if n == time {
			return AsciiView(cells)
		}
	}

	if n == time {
		return AsciiView(cells)
	}
	period = time - 3
	simulatesTodo := n % period
	for i := int32(0); i < simulatesTodo; i++ {
		simulate()
		if n == time {
			return AsciiView(cells)
		}
	}

	return AsciiView(cells)
}
