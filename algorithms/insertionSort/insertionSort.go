package insertionSort

type T interface{}

func Sort(array []T, sorter func(T, T) bool) {
	if len(array) < 2 {
		return
	}

	for i := 1; i < len(array); i++ {

		j := i
		tmp := array[j]
		for j > 0 && sorter(tmp, array[j-1]) {
			array[j] = array[j-1]
			j--
		}
		array[j] = tmp
	}
}
