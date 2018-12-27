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

	chars1 := []rune(s1)
	chars2 := []rune(s2)

	getResult := func(index1, index2 int) int32 {
		result := 0
		for index1 < len(chars1) && index2 < len(chars2) {
			ch2 := chars2[index2]

			for i := index1; i < len(chars1); i++ {
				ch1 := chars1[i]

				if ch1 == ch2 {
					result++
					index1 = i + 1
					break
				}
			}
			index2++
		}
		return int32(result)
	}

	var results []int32

	for i := 0; i < len(chars2); i++ {
		results = append(results, getResult(0, i))
	}

	return max(results)
}

func max(ar []int32) int32 {
	if len(ar) == 0 {
		return 0
	}

	max := ar[0]
	for i := 1; i < len(ar); i++ {
		if ar[i] > max {
			max = ar[i]
		}
	}
	return max
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
