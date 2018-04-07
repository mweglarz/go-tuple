package rabinKarp

import (
	"fmt"
	"math"
)

const (
	// hashMultiplier float64 = 69069
	hashMultiplier float64 = 69
)

func Search(str, pattern string) (int, error) {
	patternArray := ([]int32)(pattern)
	strArray := ([]int32)(str)

	if len(strArray) < len(patternArray) {
		return 0, fmt.Errorf("Pattern longer than text")
	}
	var (
		patternHash float64 = hash(patternArray)
		endIndx     int     = len(patternArray)
		firstChars          = strArray[:endIndx]
		firstHash           = hash(firstChars)
	)

	if patternHash == firstHash && compare(patternArray, firstChars) {
		return 0, nil
	}
	var prevHash = firstHash

	for idx := 1; idx <= len(strArray)-len(patternArray); idx++ {
		fmt.Println()
		endIndx = idx + len(patternArray)
		window := strArray[idx:endIndx]
		windowHash := nextHash(prevHash, strArray[idx-1], strArray[endIndx-1], len(patternArray)-1)

		if windowHash == patternHash && compare(patternArray, window) {
			return idx, nil
		}
		prevHash = windowHash
	}

	return 0, fmt.Errorf("Not found")
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
