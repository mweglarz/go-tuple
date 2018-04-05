package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func distribute(n int, ar []int) string {
	count := 0
	startIndex := 0
	for {
		indx, err := findTriplet(ar, startIndex)
		if err != nil {
			startIndex++
			continue
		}
		proceed := func() {
			startIndex = indx
			ar[startIndex] = ar[startIndex] + 1
			ar[startIndex+1] = ar[startIndex+1] + 2
			ar[startIndex+2] = ar[startIndex+2] + 1
			count += 4
		}
		if indx == -1 {
			break
		}
		proceed()
		startIndex += 3
	}

	if count == 0 {
		return "NO"
	}
	return strconv.Itoa(count)
}

func findTriplet(ar []int, startIndex int) (int, error) {
	if startIndex >= len(ar)-2 {
		return -1, nil
	}

	for i := startIndex; i < len(ar)-2; i++ {
		if ar[i]%2 == 1 && ar[i+1]%2 == 0 && ar[i+2]%2 == 1 {
			return i, nil
		}
	}
	return 0, errors.New("")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	arStr, _ := reader.ReadString('\n')
	nStr = strings.TrimSpace(nStr)
	arStr = strings.TrimSpace(arStr)
	n, _ := strconv.Atoi(nStr)
	arStrings := strings.Split(arStr, " ")
	var ar []int
	for _, valStr := range arStrings {
		intVal, _ := strconv.Atoi(valStr)
		ar = append(ar, intVal)
	}

	result := distribute(n, ar)
	fmt.Println(result)
}
