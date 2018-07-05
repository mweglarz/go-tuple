package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	array []rune
}

func (s *Stack) push(char rune) {
	s.array = append(s.array, char)
}

func (s *Stack) pop() (rune, bool) {
	length := len(s.array)
	if length == 0 {
		return 33, false
	}
	last := s.array[length-1]
	s.array = s.array[:length-1]
	return last, true
}

func contains(runes []rune, char rune) (bool, int) {
	for indx, aRune := range runes {
		if char == aRune {
			return true, indx
		}
	}
	return false, -1
}

func checkLine(line string) bool {
	leftBrackets := []rune{'(', '{', '['}
	rightBrackets := []rune{')', '}', ']'}
	stack := Stack{}
	for _, char := range line {
		if contain, _ := contains(leftBrackets, char); contain {
			stack.push(char)
		}
		if rightContain, rightIndex := contains(rightBrackets, char); rightContain {
			leftBracket, notEmpty := stack.pop()
			if !notEmpty {
				return false
			} else {
				_, leftIndex := contains(leftBrackets, leftBracket)
				if leftIndex != rightIndex {
					return false
				}
			}
		}
	}
	return true
}

func checkBrackets(lines []string) {
	for _, line := range lines {
		result := checkLine(line)
		if result {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	var lines []string

	for tItr := 0; tItr < int(t); tItr++ {
		expression := readLine(reader)
		lines = append(lines, expression)
	}

	checkBrackets(lines)
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
