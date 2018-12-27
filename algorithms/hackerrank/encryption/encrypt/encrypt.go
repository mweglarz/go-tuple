package encrypt

import (
	"fmt"
	"math"
)

func Encrypt(s string) string {
	s = trim(s)

	l := math.Sqrt(float64(len(s)))

	lowerBound := floor(l)
	upperBound := floor(l)

	rows, cols := adjustBounds(lowerBound, upperBound, len(s))

	chars := make([][]int32, rows)
	for i := 0; i < rows; i++ {
		chars[i] = make([]int32, cols)
	}

	for i, ch := range s {
		row := i / cols
		col := i - row*cols
		chars[row][col] = ch
	}

	printChars(chars)
	result := ""

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if ch := chars[row][col]; ch != 0 {
				result += string(ch)
			}
		}
		if col != cols-1 {
			result += " "
		}
	}

	fmt.Printf("result = %+v\n", result)

	return result
}

func trim(s string) string {
	return s
}

func floor(c float64) int {
	return int(c)
}

func printChars(chars [][]int32) {

	for _, row := range chars {
		for _, ch := range row {
			fmt.Printf("%s", string(ch))
		}
		fmt.Println()
	}
	fmt.Println()
}

func adjustBounds(lowerBound, upperBound int, len int) (int, int) {

	fmt.Printf("adjustBounds start, lower = %d, upper = %d\n", lowerBound, upperBound)

	if lowerBound*upperBound >= len {
		return lowerBound, upperBound
	}

	for lowerBound*upperBound < len {

		if upperBound-lowerBound == 1 {
			lowerBound++

		} else {
			upperBound++
		}
	}

	return lowerBound, upperBound
}
