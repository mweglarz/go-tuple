package bomberman

var time int32 = 0
var hashes map[string]int32 = make(map[string]int32)

func BomberMan(n int32, grid []string) []string {
	cells := Transform(grid)

	if n%2 == 0 {
		for i := 0; i < len(cells); i++ {
			for j := 0; j < len(cells[i]); j++ {
				cells[i][j].PlantBombIfEmpty()
			}
		}
		return AsciiView(cells)
	}

	hashLength := len(cells) * len(cells[0])
	hash := make([]rune, hashLength)
	period := n

	simulate := func() {
		hash = make([]rune, hashLength)
		time++

		for i := 0; i < len(cells); i++ {
			for j := 0; j < len(cells[i]); j++ {

				cells[i][j].TimeProgress()

				// detonate
				if time%2 == 1 && cells[i][j].State == EXPLODING {

					if i != 0 {
						cells[i-1][j].UpdateState(EMPTY)
					}

					if i != len(cells)-1 && cells[i+1][j].State != WAIT2 {
						cells[i+1][j].UpdateState(EMPTY)
					}

					if j != 0 {
						cells[i][j-1].UpdateState(EMPTY)
					}

					if j != len(cells[i])-1 && cells[i][j+1].State != WAIT2 {
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

	}

	for {
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

			if n == time {
				return AsciiView(cells)
			}

			if lastTime, ok := hashes[string(hash)]; ok {
				period = int32(time) - lastTime
				break

			} else if time%2 == 1 {
				hashes[string(hash)] = time
			}
		}

	}

	if n == time {
		return AsciiView(cells)
	}
	var simulatesTodo int

	if period == 0 {
		simulatesTodo = 10e9

	} else {
		simulatesTodo = int(n-time) % int(period)
	}

	for i := 0; i < simulatesTodo; i++ {
		simulate()
		if n == time {
			return AsciiView(cells)
		}
	}

	return AsciiView(cells)
}
