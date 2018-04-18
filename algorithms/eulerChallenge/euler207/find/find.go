package find

import (
	"errors"
	"fmt"
)

func Find(k int, array []int, low, hi int) (int, error) {
	if len(array) < 1 {
		// fmt.Println("package=  find, array < 1")
		return 0, fmt.Errorf("array is empty, %v", array)
	}
	if hi <= low {
		return 0, fmt.Errorf("value not found %d", k)
	}

	mid := (low + hi - 1) / 2
	if array[mid] > k {
		return Find(k, array, low, mid)

	} else if array[mid] < k {
		return Find(k, array, mid+1, hi)

	} else {
		return mid, nil
	}
}

func FindNearestLowerSolution(k int, array []int, low, hi int) (int, error) {

	handled, indx, err := handleInputArray(k, array, low, hi)
	if handled {
		return indx, err
	}

	handled, indx, err = handleInputSlice(k, array, low, hi)
	if handled {
		return indx, err
	}

	mid := (low + hi - 1) / 2

	if k > array[mid] && k < array[mid+1] {
		return mid, nil
	} else if k < array[mid] {
		return FindNearestLowerSolution(k, array, low, mid+1)

	} else if k > array[mid] {
		return FindNearestLowerSolution(k, array, mid, hi)

	} else {
		return mid, nil
	}
}

func handleInputArray(k int, array []int, low, hi int) (bool, int, error) {
	if len(array) == 0 {
		return true, 0, errors.New("array is empty")

	} else if len(array) == 1 {
		if k >= array[0] {
			return true, 0, nil
		} else {
			return true, 0, errors.New("not found")
		}

	} else if len(array) == 2 {
		if k >= array[1] {
			return true, 1, nil

		} else if k >= array[0] {
			return true, 0, nil

		} else {
			return true, 0, errors.New("not found")
		}
	}

	return false, 0, nil
}

func handleInputSlice(k int, array []int, low, hi int) (bool, int, error) {
	if low >= hi {
		return true, 0, errors.New("empty slice")
	}

	if len(array[low:hi]) == 1 {
		if k >= array[low] {
			return true, low, nil

		} else {
			return true, 0, errors.New("not found")
		}

	} else if len(array[low:hi]) == 2 {
		if k >= array[hi-1] {
			return true, hi - 1, nil

		} else if k >= array[low] {
			return true, low, nil

		} else {
			return true, 0, errors.New("not found")
		}
	}
	return false, 0, nil
}
