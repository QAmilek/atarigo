package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Start testing\n")
}

var stoneColors = map[string]int{
	"B": 1,
	"W": 2,
}

var coordinates = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
}

type Stone struct {
	Color int
	X     int
	Y     int
}

func buildEmptyBoard(size int) [][]int {
	board := make([][]int, size)

	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
	}

	return board
}

func (stone *Stone) putOnBoard(board [][]int) [][]int {
	board[stone.X][stone.Y] = stone.Color

	return board
}

func (stone *Stone) isMovePossible(board [][]int) bool {
	placeOnBoard := board[stone.X][stone.Y]

	if placeOnBoard != 0 {
		return false
	}

	return true
}

func playGame(moves []Stone, board [][]int) [][]int {
	for _, stone := range moves {
		if stone.isMovePossible(board) {
			stone.putOnBoard(board)
		} else {
			break
		}
	}

	return board
}

func (stone *Stone) findNeighboors(board [][]int) []Stone {
	liberties := []Stone{}

	if stone.Y > 0 {
		liberties = append(liberties, Stone{board[stone.X][stone.Y-1], stone.X, stone.Y - 1})
	}
	if stone.Y < len(board) - 1 {
		liberties = append(liberties, Stone{board[stone.X][stone.Y+1], stone.X, stone.Y + 1})
	}
	if stone.X > 0 {
		liberties = append(liberties, Stone{board[stone.X-1][stone.Y], stone.X - 1, stone.Y})
	}
	if stone.X < len(board) - 1 {
		liberties = append(liberties, Stone{board[stone.X+1][stone.Y], stone.X + 1, stone.Y})
	}

	return liberties
}

func (stone *Stone) findOpponents(board [][]int) []Stone {
	neighboors := stone.findNeighboors(board)
	opponents := []Stone{}

	for _, liberty := range neighboors {
		if liberty.Color != 0 && stone.Color != liberty.Color {
			opponents = append(opponents, liberty)
		}
	}

	return opponents
}

func findFriendsForStone(stone Stone, board [][]int) []Stone {
	neighboors := stone.findNeighboors(board)
	friends := []Stone{}

	for _, liberty := range neighboors {
		if stone.Color == liberty.Color {
			friends = append(friends, liberty)
		}
	}

	return friends
}

func isStoneInGroup(stone Stone, group []Stone) bool {
	for _, element := range group {
		if stone == element {
			return true
		}
	}

	return false
}

func makeGroupForStone(stone Stone, board [][]int) []Stone {
	group := []Stone{}
	toCheck := []Stone{}
	friends := []Stone{}
	group = append(group, stone)
	toCheck = append(toCheck, stone)

	for len(toCheck) > 0 {
		firstToCheck := toCheck[:1]

		friends = findFriendsForStone(firstToCheck[0], board)
		for _, friend := range friends {
			if !isStoneInGroup(friend, group) {
				group = append(group, friend)
				toCheck = append(toCheck, friend)
			}
		}
		toCheck = toCheck[1:]
	}

	return group
}

func findLibertiesForStone(stone Stone, board [][]int) []Stone {
	neighboors := stone.findNeighboors(board)
	liberties := []Stone{}

	for _, liberty := range neighboors {
		if liberty.Color == 0 {
			liberties = append(liberties, liberty)
		}
	}

	return liberties
}

func findLibertiesForGroup(group []Stone, board [][]int) []Stone {
	groupLiberties := []Stone{}
	stoneLiberties := []Stone{}

	for _, stone := range group {
		stoneLiberties = findLibertiesForStone(stone, board)
		for _, liberty := range stoneLiberties {
			if !isStoneInGroup(liberty, groupLiberties) {
				groupLiberties = append(groupLiberties, liberty)
			}
		}
	}

	return groupLiberties
}

func countStoneLiberties(stone Stone, board [][]int) int {
	liberties := findLibertiesForStone(stone, board)

	return len(liberties)
}

func countGroupLiberties(group []Stone, board [][]int) int {
	liberties := findLibertiesForGroup(group, board)

	return len(liberties)
}

func isGroupAlive(group []Stone, board [][]int) bool {
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

func transformMoveToStone(move string) Stone {
	color := move[0:1]
	x := move[2:3]
	y := move[3:4]

	return Stone{
		stoneColors[color],
		coordinates[x],
		coordinates[y],
	}
}

func transformMovesToStones(moves []string) []Stone {
	stones := []Stone{}
	for _, move := range moves {
		stones = append(stones, transformMoveToStone(move))
	}

	return stones
}

func getMovesFromGameRecord(record string) []Stone {
	moves := getMovesPartOfGameRecordAsArray(record)

	return transformMovesToStones(moves)
}

func writeGameOnBoard(record string) [][]int {
	board := buildEmptyBoard(3)

	return playGame(getMovesFromGameRecord(record), board)
}
