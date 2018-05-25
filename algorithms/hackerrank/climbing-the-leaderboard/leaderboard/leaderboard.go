package leaderboard

import (
	"fmt"

	search "tuple-mw.com/algorithms/binarySearch"
)

func ClimbingLeaderboard(playersScore []int32, gameScores []int32) []int32 {

	// TODO: create sorted set from playersscore, set index determine ranking
	// TODO: set have to keep track of number of players with specific score so
	// we can delete score when it's only Alice and she goes up
	var ranks []int32
	fmt.Println("climbing before prepare data")
	scoreSet := prepareData(playersScore)
	fmt.Println("climbing after prepare data")
	for _, gameScore := range gameScores {
		fmt.Println("Getting rank for score", gameScore)
		rank := getRank(scoreSet, gameScore)
		ranks = append(ranks, rank)
	}
	return ranks
}

func getRank(scoreSet []int32, gameScore int32) int32 {
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

	return int32(result) + 1
}

func prepareData(playersScore []int32) []int32 {
	var scoreSet []int32

	for _, score := range playersScore {
		valIndex, _ := search.SearchDecreasing(scoreSet, score)

		if valIndex == -1 {
			scoreSet = append(scoreSet, score)
		}
	}
	return scoreSet
}
