package repeatingVariations

import "fmt"

func FindVariations(n int, ops string, prev string) {

	if n == 1 {
		for _, char := range ops {
			fmt.Printf("%s%s\n", prev, string(char))
		}
		return
	}

	for _, char := range ops {
		FindVariations(n-1, ops, prev+string(char))
	}
}
