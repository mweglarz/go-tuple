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
	var (
		patternHash float64 = hash(patternArray)
		endIndx     int     = len(patternArray)
		firstChars          = strArray[:endIndx]
		firstHash           = hash(firstChars)
	)

	if patternHash == firstHash && compare(patternArray, firstChars) {
		return
	}
	var prevHash = firstHash

	for idx := 1; idx <= len(strArray)-len(patternArray); idx++ {
		endIndx = idx + len(patternArray)
		window := strArray[idx:endIndx]
		windowHash := nextHash(prevHash, strArray[idx-1], strArray[idx], len(patternArray)-1)

		if windowHash == patternHash && compare(patternArray, window) {
			indx = idx
			return
		}
		prevHash = windowHash
	}

	err = fmt.Errorf("Not found")
	return
}

func hash(ar []int32) (total float64) {
	var exponent int = len(ar) - 1
	for _, v := range ar {
		total += float64(v) * math.Pow(float64(hashMultiplier), float64(exponent))
		exponent--
	}
	return
}

func nextHash(prevHash float64, dropped, added int32, patternSize int) float64 {
	oldHash := prevHash - float64(dropped)*math.Pow(float64(hashMultiplier), float64(patternSize))
	return hashMultiplier*oldHash + float64(added)
}

func wordToIntArray(str string) (ar []int32) {
	for _, v := range str {
		ar = append(ar, v)
	}
	return
}

func compare(ar1 []int32, ar2 []int32) bool {
	if len(ar1) != len(ar2) {
		return false
	}

	for i := 0; i < len(ar1); i++ {
		if ar1[i] != ar2[i] {
			return false
		}
	}
	return true
}
