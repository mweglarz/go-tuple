package big

import "strings"
import "fmt"

func findMin(runes []rune, start int, validation func(rune) bool) int {
	fmt.Println("findMin invoked")
	fmt.Printf("runes = %+v\n", runes)
	fmt.Printf("start = %+v\n", start)

	minIndex := -1
	min := rune(255)
	for i := start; i < len(runes); i++ {
		if runes[i] < min && validation(runes[i]) {
			min = runes[i]
			minIndex = i
		}
	}
	fmt.Printf("minIndex = %+v\n", minIndex)
	return minIndex
}

func BiggerIsGreater(word string) string {
	fmt.Println("invoking bigger is greater, runes", []rune(word))
	runes := []rune(strings.ToLower(word))
	swap := func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	}

	swapedIndex := -1

	for i := len(runes) - 2; i >= 0; i-- {
		maxValid := func(char rune) bool {
			return char > runes[i]
		}
		if swapIndex := findMin(runes, i+1, maxValid); swapIndex != -1 {
			swapedIndex = i
			swap(i, swapIndex)
			break
		}
	}

	if swapedIndex == -1 {
		return "no answer"
	}

	for i := swapedIndex + 1; i < len(runes)-1; i++ {
		minValid := func(char rune) bool {
			return char < runes[i]
		}
		if swapIndex := findMin(runes, i+1, minValid); swapIndex != -1 {
			swap(i, swapIndex)
		}
	}

	fmt.Printf("final runes = %+v\n", runes)
	return string(runes)
}
