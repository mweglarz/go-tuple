package container

func OrganizingContainers(container [][]int32) bool {

	size := len(container)

	containerMap := make(map[int]int32)
	typesMap := make(map[int]int32)

	for i := 0; i < size; i++ {

		for j := 0; j < size; j++ {
			containerMap[i] = containerMap[i] + container[i][j]
			typesMap[j] = typesMap[j] + container[i][j]
		}
	}

	for _, count := range containerMap {
		found := false

		for ballType, ballCount := range typesMap {
			if count == ballCount {
				found = true
				typesMap[ballType] = -1
				break
			}
		}

		if !found {
			return false
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
