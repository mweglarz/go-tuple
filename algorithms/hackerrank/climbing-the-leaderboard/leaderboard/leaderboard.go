package leaderboard

import (
	"fmt"

	search "tuple-mw.com/algorithms/binarySearch"
)

func GetRankProgress(playersCount int, playersScore []int, gameCount int, gameScores []int) []int {

	// TODO: create sorted set from playersscore, set index determine ranking
	// TODO: set have to keep track of number of players with specific score so
	// we can delete score when it's only Alice and she goes up
	var ranks []int
	scoreSet, _ := prepareData(playersScore)
	for _, gameScore := range gameScores {
		fmt.Println("Getting rank for score", gameScore)
		rank := getRank(scoreSet, gameScore)
		ranks = append(ranks, rank)
	}
	return ranks
}

func getRank(scoreSet []int, gameScore int) int {
	fmt.Println("scoreSet", scoreSet, "gameScore", gameScore)
	index, nearestIndex := search.SearchDecreasing(scoreSet, gameScore)
	fmt.Printf("getRank, index = %d, nearestIndex = %d\n", index, nearestIndex)

	var result int

	if index != -1 {
		result = index

	} else if nearestIndex == -1 {
		result = 0

	} else {
		nearestValue := scoreSet[nearestIndex]
		if nearestValue > gameScore {
			result = nearestIndex + 1

		} else {
			result = nearestIndex
		}
	}

	return result + 1
}

func prepareData(playersScore []int) ([]int, []int) {
	var scoreSet []int
	var scoreCount []int

	for _, score := range playersScore {
		valIndex, _ := search.SearchDecreasing(scoreSet, score)

		if valIndex == -1 {
			scoreSet = append(scoreSet, score)
			// scoreCount = append(scoreCount, 1)

		} //  else {
		// 	fmt.Printf("%d found at index %d incrementing %d\n", score, valIndex, scoreCount[valIndex])
		// 	// scoreCount[valIndex] += 1
		// }
	}
	return scoreSet, scoreCount
}
