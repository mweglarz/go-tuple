package rabinKarp

import (
	"fmt"
	"math"
)

const (
	// hashMultiplier float64 = 69069
	hashMultiplier float64 = 69
)

func Search(str, pattern string) (indx int, err error) {
	fmt.Printf("str = %+v\n", str)
	fmt.Printf("pattern = %+v\n", pattern)

	patternArray := ([]int32)(pattern)
	strArray := ([]int32)(str)
	// fmt.Printf("patternArray = %+v\n", patternArray)
	// fmt.Printf("strArray = %+v\n", strArray)

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
	fmt.Printf("patternHash = %+v\n", patternHash)
	fmt.Printf("firstChars = %+v, %v\n", firstChars, string(firstChars))
	fmt.Printf("firstHash = %+v\n", firstHash)

	if patternHash == firstHash && compare(patternArray, firstChars) {
		return
	}
	var prevHash = firstHash

	for idx := 0; idx < len(strArray)-len(patternArray); idx++ {
		fmt.Println()
		endIndx = idx + len(patternArray)
		window := strArray[idx:endIndx]
		fmt.Printf("window = %+v, %v\n", window, string(window))
		fmt.Printf("prevHash = %+v\n", prevHash)
		fmt.Printf("idx = %+v\n", idx)
		fmt.Printf("strArray[idx] = %+v\n", strArray[idx])
		fmt.Printf("strArray[endIndx] = %+v\n", strArray[endIndx])

		windowHash := nextHash(prevHash, strArray[idx], strArray[endIndx], len(patternArray)-1)
		fmt.Printf("windowHash = %+v\n", windowHash)

		if windowHash == patternHash {
			fmt.Println("window hash match pattern hash")
			fmt.Printf("patternArray = %+v\n", patternArray)
			fmt.Printf("window = %+v\n", window)
			if compare(patternArray, window) {
				fmt.Println("window match pattern, returning index", idx)
				indx = idx
				return
			}
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
