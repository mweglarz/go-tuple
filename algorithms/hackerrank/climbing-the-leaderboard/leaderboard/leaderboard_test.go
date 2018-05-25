package leaderboard

import (
	"fmt"
	"testing"
)

var input []int32 = []int32{100, 100, 50, 40, 40, 20, 10}
var expectScore []int32 = []int32{100, 50, 40, 20, 10}

func TestRankForSampleInput(t *testing.T) {

	var (
		playersScore []int32 = []int32{100, 100, 50, 40, 40, 20, 10}
		gameScores   []int32 = []int32{5, 25, 50, 120}
		expect       []int32 = []int32{6, 4, 2, 1}
	)

	result := ClimbingLeaderboard(playersScore, gameScores)
	fmt.Println("rankForSampleInput result", result)

	for i := 0; i < len(gameScores); i++ {
		if result[i] != expect[i] {
			t.Fatalf("error at position %d, expected %d, got %d", i, expect[i], result[i])
		}
	}
}

func TestPrepareData(t *testing.T) {

	outScore := prepareData(input)

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
	var gameScore int32 = 50
	var expectPosition int32 = 2

	scoreSet := prepareData(input)
	calcPosition := getRank(scoreSet, gameScore)
	if calcPosition != expectPosition {
		t.Fatalf("expected position %d differ from calculated position %d, score %d, scoreSet %v", expectPosition, calcPosition, gameScore, scoreSet)
	}
}

func TestGetRank2(t *testing.T) {
	var gameScore int32 = 60
	var expectPosition int32 = 2

	scoreSet := prepareData(input)
	calcPosition := getRank(scoreSet, gameScore)
	if calcPosition != expectPosition {
		t.Fatalf("expected position %d differ from calculated position %d", expectPosition, calcPosition)
	}
}

func TestGetRank3(t *testing.T) {
	var gameScore int32 = 5
	var expectPosition int32 = 6

	scoreSet := prepareData(input)
	calcPosition := getRank(scoreSet, gameScore)
	if calcPosition != expectPosition {
		t.Fatalf("expected position %d differ from calculated position %d", expectPosition, calcPosition)
	}
}
