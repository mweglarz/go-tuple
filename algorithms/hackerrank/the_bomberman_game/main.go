package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	_ "strconv"
	"strings"
	"tuple-mw.com/algorithms/hackerrank/the_bomberman_game/bomberman"
)

// Complete the bomberMan function below.
func bomberMan(n int32, grid []string) []string {
	return bomberman.BomberMan(n, grid)
}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout := os.Stdout
	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	// rcn := strings.Split(readLine(reader), " ")

	// rTemp, err := strconv.ParseInt(rcn[0], 10, 64)
	// checkError(err)
	// r := int32(rTemp)

	// cTemp, err := strconv.ParseInt(rcn[1], 10, 64)
	// checkError(err)
	// c := int32(cTemp)
	// _ = c

	// nTemp, err := strconv.ParseInt(rcn[2], 10, 64)
	// checkError(err)
	// n := int32(nTemp)
	n := int32(3)
	grid := []string{
		".......",
		"...O...",
		"....O..",
		".......",
		"OO.....",
		"OO.....",
	}

	// var grid []string

	// for i := 0; i < int(r); i++ {
	// gridItem := readLine(reader)
	// grid = append(grid, gridItem)
	// }

	result := bomberMan(n, grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
