package boyerMooreStringSearch

import "errors"

func IndexOf(pattern string, inString string) (int, error) {
	var patternLength int = len(pattern)

	if !(patternLength > 0 && patternLength < len(inString)) {
		return 0, errors.New("Wrong pattern length")
	}

	var stringLength = len(inString)
	var skipTable map[rune]int = make(map[rune]int)
	for i, ch := range pattern {
		skipTable[ch] = patternLength - i - 1
	}

	patternEndIndex := patternLength - 1
	lastChar := rune(pattern[patternEndIndex])
	var i int = patternEndIndex

	backwards := func() (int, error) {
		q := patternEndIndex
		j := i
		for q > 0 {
			j, q = j-1, q-1
			if inString[j] != pattern[q] {
				return 0, errors.New("Not found")
			}
		}
		return j, nil
	}

	jump := func(c rune, idx int) int {
		if offset, ok := skipTable[c]; ok {
			idx += offset
		} else {
			idx += patternLength
		}
		if idx >= stringLength {
			idx = stringLength - 1
		}
		if idx == i {
			idx++
		}
		return idx
	}

	var c rune
	for ; i < stringLength; i = jump(c, i) {
		c = rune(inString[i])

		if c == lastChar {
			if k, err := backwards(); err == nil {
				return k, nil
			}
		}
	}
	return 0, errors.New("Pattern not found")
}
