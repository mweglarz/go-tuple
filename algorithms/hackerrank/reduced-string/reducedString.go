package reducedString

import (
	"errors"
)

func Reduce(str string) string {

	var err error
	var tempStr string = str
	for err == nil {
		tempStr, err = reduceOnce(tempStr)
	}
	if len(tempStr) == 0 {
		return "Empty String"
	} else {
		return tempStr
	}
}

func reduceOnce(str string) (string, error) {
	if len(str) < 2 {
		return "", errors.New("string len less than 2")
	}
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			if i+2 < len(str) {
				return str[:i] + str[i+2:], nil
			} else {
				return str[:i], nil
			}
		}
	}

	return str, errors.New("no more duplicates")
}
