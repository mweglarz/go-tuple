package container

func OrganizingContainers(container [][]int32) bool {

	size := len(container)

	// iterate types
	for i := 0; i < size; i++ {
		// iterate containers
		for j := 0; j < size; j++ {
			if i == j {
				continue
			}

			min := min(container[i][j], container[j][i])

		}
	}

	return false
}

func min(a, b int32) int32 {
	min := a
	if b < a {
		min = b
	}
	return min
}
