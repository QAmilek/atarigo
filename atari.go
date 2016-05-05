package main

//import (
//	"fmt"
//)

func main () {
}

type Stone struct {
	Color int
	X int
	Y int
}

func buildEmptyBoard (size int) [][]int {
	board := make([][]int, size)

	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
	}

	return board
}

func putStoneOnBoard(stone Stone, board [][]int) [][]int {
	board[stone.X][stone.Y] = stone.Color

	return board
}

func assertIllegalMove(stone Stone, board [][]int) bool {
	if check := board[stone.X][stone.Y]; check == 0  {
		return true
	}

	if check := board[stone.X][stone.Y]; check == stone.Color  {
		//fmt.Println("Your stone is already there")
		return false
	}

	if check := board[stone.X][stone.Y]; check != stone.Color {
		//fmt.Println("This place is occupied by your opponent stone")
		return false
	}

	return true
}

func playGame(moves []Stone, board [][]int) [][]int {
	for _, stone := range moves {
		if assertIllegalMove(stone, board) {
			putStoneOnBoard(stone, board)
		} else {
			break
		}
	}

	return board
}

func findNeighboors(stone Stone, board [][]int) []Stone {
	liberties := []Stone{}

	if limit := stone.Y; limit > 0 {
		liberties = append(liberties, Stone{board[stone.X][stone.Y - 1], stone.X, stone.Y - 1})
	}
	if limit := stone.Y; limit < cap(board) - 1 {
		liberties = append(liberties, Stone{board[stone.X][stone.Y + 1], stone.X, stone.Y + 1})
	}
	if limit := stone.X; limit > 0 {
		liberties = append(liberties, Stone{board[stone.X - 1][stone.Y], stone.X - 1, stone.Y})
	}
	if limit := stone.X; limit < cap(board) - 1 {
		liberties = append(liberties, Stone{board[stone.X + 1][stone.Y], stone.X + 1, stone.Y})
	}

	return liberties
}

func findOpponentForStone(stone Stone, board [][]int) []Stone {
	neighboors := findNeighboors(stone, board)
	opponents := []Stone{}

	for _, liberty := range neighboors {
		if liberty.Color != 0 && stone.Color != liberty.Color {
			opponents = append(opponents, liberty)
		}
	}

	return opponents
}

func findFriendsForStone(stone Stone, board [][]int) []Stone {
	neighboors := findNeighboors(stone, board)
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
	group = append(group, stone)
	friends := findFriendsForStone(stone, board);

	for len(friends) > 0 {
		for _, friend := range friends {
			group = append(group, friend)
		}

		friends = friends[1:]
	}

	return group
}
