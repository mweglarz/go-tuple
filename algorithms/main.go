package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	num, _ := strconv.Atoi(numStr)
	var matrix [][]int8
	for i := 0; i < num; i++ {
		rowString, _ := reader.ReadString('\n')
		rowString = strings.TrimSpace(rowString)
		matrix = append(matrix, stringToRow(rowString))
	}
	strMatrix := matrixToStringMatrix(num, matrix)
	find2DPeak(num, matrix, strMatrix)
	printMatrix(strMatrix)
}

func find2DPeak(n int, matrix [][]int8, strMatrix [][]string) {
	if n < 3 {
		return
	}

	for i := 1; i < n-1; i++ {
		peaks := find1DPeak3(matrix[i])
		if len(peaks) < 1 {
			continue
		}
		// j, err := find1DPeak(matrix[i], 0)
		// fmt.Println("1dpeak result", j, err)
		// if err != nil || j == 0 || j == n-1 {
		// 	continue
		// }
		// fmt.Printf("checking if %d,%d is peak\n", i, j)
		for _, j := range peaks {
			if is2DPeak(matrix, i, j) {
				// fmt.Println("peak found!")
				strMatrix[i][j] = "X"
			}
		}
	}
}

func is2DPeak(ar [][]int8, i, j int) bool {
	// fmt.Println("checking i", i, "j", j)
	val := ar[i][j]
	// fmt.Printf("ar[i][j] = %+v\n", ar[i][j])
	// fmt.Printf("ar[i+1][j] = %+v\n", ar[i+1][j])
	// fmt.Printf("ar[i-1][j] = %+v\n", ar[i-1][j])
	// fmt.Printf("ar[i][j-1] = %+v\n", ar[i][j-1])
	// fmt.Printf("ar[i][j+1] = %+v\n", ar[i][j+1])

	if ar[i+1][j] < val && ar[i-1][j] < val && ar[i][j-1] < val && ar[i][j+1] < val {
		return true
	}
	return false
}

func find1DPeak3(ar []int8) []int {
	prev := ar[0]
	next := ar[0]
	var peaks []int

	for i := 1; i < len(ar)-1; i++ {
		prev = ar[i-1]
		next = ar[i+1]
		if ar[i] > prev && ar[i] > next {
			peaks = append(peaks, i)
		}
	}
	return peaks
}

func find1DPeak2(ar []int8, low, hi int) int {
	if low == hi {
		return low
	}
	if low == hi-1 {
		indx := low
		if ar[low] < ar[hi] {
			indx = hi
		}
		return indx
	}

	mid := (low + hi) / 2
	if ar[mid] >= ar[mid+1] {
		return find1DPeak2(ar, low, mid)
	} else {
		return find1DPeak2(ar, mid+1, hi)
	}
}

func find1DPeak(ar []int8, offset int) (int, error) {
	// fmt.Println("find1DPeak")
	// fmt.Println(ar)
	count := len(ar)

	if count == 2 {
		indx, _ := max(ar[0], ar[1])
		return offset + indx, nil
	} else if count < 3 {
		return 0, fmt.Errorf("not found")
	}

	mid := count / 2
	if ar[mid] < ar[mid-1] {
		// fmt.Println("left", ar[mid], ar[mid-1])
		return find1DPeak(ar[:mid], offset)
	} else if ar[mid] < ar[mid+1] {
		// fmt.Println("right", ar[mid], ar[mid+1])
		return find1DPeak(ar[mid:], offset+mid)
	} else {
		return offset + mid, nil
	}
	// return 0, fmt.Errorf("not found")
}

func matrixToStringMatrix(n int, matrix [][]int8) [][]string {
	// fmt.Println(matrix)
	strMatrix := make([][]string, n)
	for i := 0; i < n; i++ {
		strMatrix[i] = make([]string, n)
		for j := 0; j < n; j++ {
			intValue := int(matrix[i][j])
			strMatrix[i][j] = strconv.Itoa(intValue)
		}
	}
	return strMatrix
}

func stringToRow(line string) (ar []int8) {
	for _, char := range line {
		num, _ := strconv.Atoi(string(char))
		ar = append(ar, int8(num))
	}
	return
}

func printMatrix(m [][]string) {
	for _, row := range m {
		for _, value := range row {
			fmt.Printf("%s", value)
		}
		fmt.Println()
	}
}

func max(a, b int8) (int, int8) {
	result := a
	indx := 0
	if b > a {
		indx = 1
		result = b
	}
	return indx, result
}
