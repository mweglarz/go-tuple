package leaderboard

import "testing"

var input []int = []int{100, 100, 50, 40, 40, 20, 10}
var expectScore []int = []int{100, 50, 40, 20, 10}
var expectCount []int = []int{2, 1, 2, 1, 1}

func TestRankForSampleInput(t *testing.T) {

	var (
		playersCount int   = 7
		playersScore []int = []int{100, 100, 50, 40, 40, 20, 10}
		gameCount    int   = 4
		gameScores   []int = []int{5, 25, 50, 120}
		expect       []int = []int{6, 4, 2, 1}
	)

	result := GetRankProgress(playersCount, playersScore, gameCount, gameScores)

	for i := 0; i < gameCount; i++ {
		if result[i] != expect[i] {
			t.Fatalf("error at position %d, expected %d, got %d", i, expect[i], result[i])
		}
	}
}

func TestPrepareData(t *testing.T) {

	outScore, outCount := prepareData(input)

	if len(expectScore) != len(outScore) || len(expectCount) != len(outCount) {
		t.Fatalf("Array sizes don't match, expectScore %d, outScore %d, expectCount %d, outCount %d", len(expectScore), len(outScore), len(expectCount), len(outCount))
	}

	for i := 0; i < len(expectScore); i++ {
		if expectScore[i] != outScore[i] || expectCount[i] != outCount[i] {
			t.Fatal("Expect and out don't match")
		}
	}
}

func TestGetRank(t *testing.T) {
	gameScore := 50
	expectPosition := 2

	calcPosition := getRank(input, gameScore) + 1 // array counts from 0
	if calcPosition != expectPosition {
		t.Fatalf("expected position %d differ from calculated position %d", expectPosition, calcPosition)
	}
}

func TestGetRank2(t *testing.T) {
	gameScore := 60
	expectPosition := 2

	calcPosition := getRank(input, gameScore) + 1 // array counts from 0
	if calcPosition != expectPosition {
		t.Fatalf("expected position %d differ from calculated position %d", expectPosition, calcPosition)
	}
}

func TestGetRank3(t *testing.T) {
	gameScore := 5
	expectPosition := 6

	calcPosition := getRank(input, gameScore) + 1 // array counts from 0
	if calcPosition != expectPosition {
		t.Fatalf("expected position %d differ from calculated position %d", expectPosition, calcPosition)
	}
}
