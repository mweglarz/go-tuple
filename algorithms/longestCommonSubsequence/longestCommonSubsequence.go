package longestCommonSubsequence

import (
	"fmt"
)

func LCS(s1, s2 string) (result string, err error) {
	matrix := lcsLength(s1, s2)
	result = backtrack(matrix, s1, s2)
	return
}

func lcsLength(s1, s2 string) [][]int {
	var matrix [][]int = repeating(len(s1)+1, len(s2)+1, 0)
	// printMatrix(matrix)
	for i, ch1 := range s1 {
		for j, ch2 := range s2 {
			if ch1 == ch2 {
				matrix[i+1][j+1] = matrix[i][j] + 1

			} else {
				matrix[i+1][j+1] = max(matrix[i+1][j], matrix[i][j+1])
			}
		}
	}
	// fmt.Println()
	// printMatrix(matrix)
	return matrix
}

func backtrack(matrix [][]int, s1, s2 string) string {
	var indexInSequence = len(s1)
	var lcs []rune

	for i, j := len(s1), len(s2); i >= 1 && j >= 1; {
		current := matrix[i][j]
		if current == matrix[i][j-1] {
			j--

		} else if current == matrix[i-1][j] {
			i--
			indexInSequence--

		} else {
			i--
			j--
			indexInSequence--
			lcs = append(lcs, rune(s1[indexInSequence]))
		}
	}
	reverse(lcs)
	return string(lcs)
}

func repeating(n, m, value int) (matrix [][]int) {
	for i := 0; i < n; i++ {
		var row []int
		for j := 0; j < m; j++ {
			row = append(row, value)
		}
		matrix = append(matrix, row)
	}
	return
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d, ", val)
		}
		fmt.Println()
	}
}

func max(a, b int) (val int) {
	val = a
	if b > a {
		val = b
	}
	return
}

func reverse(ar []rune) {
	for i, j := 0, len(ar)-1; i < j; i, j = i+1, j-1 {
		ar[i], ar[j] = ar[j], ar[i]
	}
}
