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

func putStonesOnBoard(moves Stones, board [][]int) [][]int {
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
	for i := 3; i < len(board); i++ {
		xAxis = append(xAxis, "  ")
	}

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
	r := regexp.MustCompile(`SZ\[(\d+)\]`)
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
	board := getEmptyBoardFromRecord(record)
	moves := getMovesFromGameRecord(record)

	return putStonesOnBoard(moves, board)
}

func getEmptyBoardFromRecord(record string) [][]int {
	size := getBoardSizeFromRecord(record)

	return buildEmptyBoard(size)
}

func markDeadGroup(group Stones, board [][]int) [][]int {
       for _, stone := range group {
               board[stone.X][stone.Y] = stone.Color + 2
       }

       return board
}

func playGame(record string) string {
       board := getEmptyBoardFromRecord(record)
       result := ""
       moves := getMovesFromGameRecord(record)

       for _, move := range moves {
               board = move.PutOnBoard(board)
               if move.Color == 1 {
                       result = "White to play"
               } else {
                       result = "Black to play"
               }

               opponents := move.FindOpponents(board)
               if len(opponents) == 0 {
                       continue
               }

               for _, opponent := range opponents {
                       group := opponent.MakeGroup(board)
                       groupStatus := isGroupAlive(group, board)

                       if !groupStatus && move.Color == 1 {
                               board = markDeadGroup(group, board)
                               return "Black won"
                       }

                       if !groupStatus && move.Color == 2 {
                               board = markDeadGroup(group, board)
                               return "White won"
                       }
               }
       }
       // printBoardToConsole(board)
       return result
}
