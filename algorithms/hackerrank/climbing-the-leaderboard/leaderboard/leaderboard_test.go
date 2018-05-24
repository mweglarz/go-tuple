package leaderboard

import (
	"fmt"
	"testing"
)

var input []int = []int{100, 100, 50, 40, 40, 20, 10}
var expectScore []int = []int{100, 50, 40, 20, 10}

func TestRankForSampleInput(t *testing.T) {

	var (
		playersCount int   = 7
		playersScore []int = []int{100, 100, 50, 40, 40, 20, 10}
		gameCount    int   = 4
		gameScores   []int = []int{5, 25, 50, 120}
		expect       []int = []int{6, 4, 2, 1}
	)

	result := GetRankProgress(playersCount, playersScore, gameCount, gameScores)
	fmt.Println("rankForSampleInput result", result)

	for i := 0; i < gameCount; i++ {
		if result[i] != expect[i] {
			t.Fatalf("error at position %d, expected %d, got %d", i, expect[i], result[i])
		}
	}
}

func TestPrepareData(t *testing.T) {

	outScore, _ := prepareData(input)

	if len(expectScore) != len(outScore) {
		t.Fatalf("Array sizes don't match, expectScore %d, outScore %d", len(expectScore), len(outScore))
	}

	for i := 0; i < len(expectScore); i++ {
		if expectScore[i] != outScore[i] {
			t.Fatal("Expect and out don't match")
		}
	}
}

func TestGetRank(t *testing.T) {
	gameScore := 50
	expectPosition := 2

	scoreSet, _ := prepareData(input)
	calcPosition := getRank(scoreSet, gameScore)
	if calcPosition != expectPosition {
		t.Fatalf("expected position %d differ from calculated position %d, score %d, scoreSet %v", expectPosition, calcPosition, gameScore, scoreSet)
	}
}

func TestGetRank2(t *testing.T) {
	gameScore := 60
	expectPosition := 2

	scoreSet, _ := prepareData(input)
	calcPosition := getRank(scoreSet, gameScore)
	if calcPosition != expectPosition {
		t.Fatalf("expected position %d differ from calculated position %d", expectPosition, calcPosition)
	}
}

func TestGetRank3(t *testing.T) {
	gameScore := 5
	expectPosition := 6

	scoreSet, _ := prepareData(input)
	calcPosition := getRank(scoreSet, gameScore)
	if calcPosition != expectPosition {
		t.Fatalf("expected position %d differ from calculated position %d", expectPosition, calcPosition)
	}
}
