package rabinKarp

import (
	"fmt"
	"math"
)

const (
	hashMultiplier float64 = 69069
)

func Search(str, pattern string) (indx int, err error) {

	patternArray := wordToIntArray(pattern)
	strArray := wordToIntArray(str)

	if len(strArray) < len(patternArray) {
		err = fmt.Errorf("Pattern longer than text")
		return
	}

	return
}

func hash(ar []int) (total float64) {
	var exponent int = len(ar) - 1
	for _, v := range ar {
		total += float64(v) * math.Pow(float64(hashMultiplier), float64(exponent))
		exponent--
	}
	return
}

func nextHash(prevHash float64, dropped, added, patternSize int) float64 {
	oldHash := prevHash - float64(dropped)*math.Pow(float64(hashMultiplier), float64(patternSize))
	return hashMultiplier*oldHash + float64(added)
}

func wordToIntArray(str string) (ar []int32) {
	for _, v := range str {
		ar = append(ar, v)
	}
	return
}
