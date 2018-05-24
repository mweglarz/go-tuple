package leaderboard

import (
	"fmt"

	search "tuple-mw.com/algorithms/binarySearch"
)

type Score struct {
	Value        int
	PlayersCount int
}

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
	index, nearestIndex := search.SearchDecreasing(scoreSet, gameScore)
	fmt.Printf("getRank, index = %+v\n", index)

	if index != -1 {
		return index

	} else if nearestIndex == -1 {
		return 1

	} else {
		nearestValue := scoreSet[nearestIndex]
		if nearestValue > gameScore {
			return nearestIndex + 1
		}
		return nearestIndex
	}
}

func prepareData(playersScore []int) ([]int, []int) {
	var scoreSet []int
	var scoreCount []int

	fmt.Println("prepareData BEGIN", playersScore)

	for _, score := range playersScore {
		valIndex, _ := search.SearchDecreasing(scoreSet, score)

		if valIndex == -1 {
			fmt.Printf("%d not found, inserting at index %d\n", score, len(scoreSet))
			scoreSet = append(scoreSet, score)
			// scoreCount = append(scoreCount, 1)

		} //  else {
		// 	fmt.Printf("%d found at index %d incrementing %d\n", score, valIndex, scoreCount[valIndex])
		// 	// scoreCount[valIndex] += 1
		// }

		fmt.Println("for continue, score handled", score)
	}

	fmt.Println("prepareData END")

	return scoreSet, scoreCount
}
