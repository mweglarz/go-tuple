package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the commonChild function below.
func commonChild(s1 string, s2 string) int32 {

	lcsMatrix := lcsLength(s1, s2)
	return int32(lcsMatrix[len(s1)][len(s2)])
}

func lcsLength(s1, s2 string) [][]int {
	var matrix [][]int = repeating(len(s1)+1, len(s2)+1, 0)
	for i, ch1 := range s1 {
		for j, ch2 := range s2 {
			if ch1 == ch2 {
				matrix[i+1][j+1] = matrix[i][j] + 1

			} else {
				matrix[i+1][j+1] = max(matrix[i+1][j], matrix[i][j+1])
			}
		}
	}
	return matrix
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

func max(a, b int) (val int) {
	val = a
	if b > a {
		val = b
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	stdout := os.Stdout

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s1 := readLine(reader)

	s2 := readLine(reader)

	result := commonChild(s1, s2)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
