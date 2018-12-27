package encrypt

import (
	"fmt"
	"math"
)

func Encrypt(s string) string {
	s = trim(s)

	l := math.Sqrt(float64(len(s)))

	lowerBound := floor(l)
	upperBound := ceil(l)

	rows := lowerBound
	cols := upperBound

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

func ceil(c float64) int {

	return int(c) + 1
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
