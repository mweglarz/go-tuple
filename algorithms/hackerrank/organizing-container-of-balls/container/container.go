package container

func OrganizingContainers(container [][]int32) bool {

	size := len(container)

	// iterate containers
	for i := 0; i < size; i++ {
		// iterate types
		for j := 0; j < size; j++ {
			if i == j || container[i][j] == 0 {
				continue
			}

			for l := j; l < size && container[i][j] != 0; l++ {
				countSwapped := swap(container, i, j, l, i)
				if l != j {
					container[l][j] = container[l][j] + countSwapped
				}
			}

			if container[i][j] != 0 {
				return false
			}
		}
	}

	return true
}

func swap(container [][]int32, i1, j1, i2, j2 int) int32 {
	min := min(container[i1][j1], container[i2][j2])
	container[i1][j1] = container[i1][j1] - min
	container[i2][j2] = container[i2][j2] - min
	return min
}

func min(a, b int32) int32 {
	min := a
	if b < a {
		min = b
	}
	return min
}
