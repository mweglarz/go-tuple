package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"tuple-mw.com/algorithms/hackerrank/climbing-the-leaderboard/leaderboard"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("output.dat")
	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	fmt.Println("setting output")
	checkError(err)

	defer stdout.Close()

	fmt.Println("flag 1")
	writer := bufio.NewWriterSize(stdout, 1024*1024)
	fmt.Println("flag 2")

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	fmt.Println("flag 3")

	scoresTemp := strings.Split(readLine(reader), " ")
	fmt.Println("scoresTemp len = ", len(scoresTemp))
	fmt.Println("scoresCount = ", scoresCount)
	if int64(len(scoresTemp)) != scoresCount {
		panic("scoresTemp doesn't match scoresCount")
	}

	var scores []int32

	fmt.Println("flag 4")
	for i := 0; i < int(scoresCount); i++ {
		fmt.Println("flag 5", i)
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	fmt.Println("flag 6")
	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	fmt.Println("flag 7")

	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int32

	fmt.Println("after parsing input")

	for i := 0; i < int(aliceCount); i++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
		checkError(err)
		aliceItem := int32(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	fmt.Println("calling climbing leaderboard")
	result := leaderboard.ClimbingLeaderboard(scores, alice)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
