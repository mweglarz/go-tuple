package stIO

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func read() {
	reader := bufio.NewReader(os.Stdin)
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	num, _ := strconv.Atoi(numStr)
	var matrix [][]int8
	for i := 0; i < num; i++ {
		rowString, _ := reader.ReadString('\n')
		rowString = strings.TrimSpace(rowString)
		strings.Split
		matrix = append(matrix, stringToRow(rowString))
	}
}
