package main

import (
	"fmt"
	"strings"
	"github.com/QAmilek/atarigo/stone"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Start testing\n")
}

type Stones []stone.Stone

func buildEmptyBoard(size int) [][]int {
	board := make([][]int, size)

	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
	}

	return board
}

func playGame(moves Stones, board [][]int) [][]int {
	for _, s := range moves {
		if s.IsMovePossible(board) {
			s.PutOnBoard(board)
		} else {
			break
		}
	}

	return board
}

func findLibertiesForGroup(group Stones, board [][]int) Stones {
	groupLiberties := Stones{}
	stoneLiberties := Stones{}

	for _, s := range group {
		stoneLiberties = s.FindLiberties(board)
		for _, liberty := range stoneLiberties {
			if !liberty.IsInGroup(groupLiberties) {
				groupLiberties = append(groupLiberties, liberty)
			}
		}
	}

	return groupLiberties
}

func countGroupLiberties(group Stones, board [][]int) int {
	liberties := findLibertiesForGroup(group, board)

	return len(liberties)
}

func isGroupAlive(group Stones, board [][]int) bool {
	libertiesNumber := countGroupLiberties(group, board)

	return libertiesNumber != 0
}

func printBoardToConsole(board [][]int) {
	xAxis := []string{"X ", "| ", "V "}
	fmt.Println("\n  Y->")
	for _, row := range board {
		fmt.Printf(xAxis[0])
		for _, column := range row {
			fmt.Printf("%v", column)
		}
		fmt.Printf("\n")
		xAxis = xAxis[1:]
	}
}

func getMovesPartOfGameRecordAsArray(record string) []string {
	record = strings.Trim(record, "(;")
	record = strings.Trim(record, ")")
	recordArray := strings.Split(record, ";")
	movesArray := recordArray[1:]

	return movesArray
}

func getBoardSizeFromRecord(record string) int {
	r := regexp.MustCompile(`SZ\[(?P<size>\d+)\]`)
	result := r.FindStringSubmatch(record)
	size, _ := strconv.Atoi(result[1])

	return size
}

func transformMovesToStones(moves []string) Stones {
	stones := Stones{}
	for _, move := range moves {
		stones = append(stones, stone.TransformMoveToStone(move))
	}

	return stones
}

func getMovesFromGameRecord(record string) Stones {
	moves := getMovesPartOfGameRecordAsArray(record)

	return transformMovesToStones(moves)
}

func writeGameOnBoard(record string) [][]int {
	board := buildEmptyBoard(3)

	return playGame(getMovesFromGameRecord(record), board)
}
