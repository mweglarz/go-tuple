package bomberman

type BombState uint8

const (
	PLANTED BombState = iota
	WAIT1
	WAIT2
	EXPLODING
	EMPTY
)

var time int32 = 0
var hashes map[string]int32 = make(map[string]int32)
var period int32

func Detonate(cells [][]Cell) {
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {

			if cells[i][j].State == EXPLODING {

				if i != 0 && cells[i-1][j].State != EXPLODING {
					cells[i-1][j].UpdateState(EMPTY)
				}

				if i != len(cells)-1 && cells[i+1][j].State != EXPLODING {
					cells[i+1][j].UpdateState(EMPTY)
				}

				if j != 0 && cells[i][j-1].State != EXPLODING {
					cells[i][j-1].UpdateState(EMPTY)
				}

				if j != len(cells[i])-1 && cells[i][j+1].State != EXPLODING {
					cells[i][j+1].UpdateState(EMPTY)
				}
				cells[i][j].UpdateState(EMPTY)
			}
		}
	}
}

func BomberMan(n int32, grid []string) []string {
	cells := Transform(grid)
	AdvanceToInitState(cells)
	if n == time {
		return AsciiView(cells)
	}

	hashExist := func() bool {
		_, ok := hashes[Hash(cells)]
		return ok
	}

	simulate := func() {
		time++
		ForEach(cells, func(c *Cell) {
			c.TimeProgress()
		})

		if time%2 == 1 {
			Detonate(cells)

		} else {
			ForEach(cells, func(c *Cell) {
				c.PlantBombIfEmpty()
			})
		}

	}

	for !hashExist() {
		// after detonating
		if time%2 == 1 {
			hashes[Hash(cells)] = time
		}
		simulate()
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
	}

	return AsciiView(cells)
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
