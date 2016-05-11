package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestBuildingEmptyBoard(t *testing.T) {
	cases := []struct {
		inSize int
		want [][]int
	}{
		{3, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}},
	}
	for _, c := range cases {
		got := buildEmptyBoard(c.inSize)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v", c.want, got)
		} else {
			//fmt.Println("\nTest buildEmptyBoard(boardSize) function:")
			//fmt.Printf("size: %v\nresult: %v\n", c.inSize, got)
		}
	}
}

func TestPlayingGame(t *testing.T) {
	cases := []struct {
		stones Stones
		inBoard [][]int
		want [][]int
	}{
		{Stones{{1, 1, 1}, {2, 1, 2}}, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 2}, []int{0, 0, 0}}},
		{Stones{{2, 1, 1}}, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 2, 0}, []int{0, 0, 0}}},
	}
	for _, c := range cases {
		got := playGame(c.stones, c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest playGame(stones, emptyBoard) function:")
			//fmt.Printf("stones: %v\nresult: %v\n", c.stones, c.want)
		}
	}
}

func TestPlayingGameWithIllegalMove(t *testing.T) {
	cases := []struct {
		stones Stones
		inBoard [][]int
		want [][]int
	}{
		{Stones{{1, 1, 1}, {2, 1, 2}, {1, 1, 2}}, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 2}, []int{0, 0, 0}}},
		{Stones{{1, 1, 1}, {2, 1, 2}, {1, 1, 1}, {2, 2, 2}}, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 2}, []int{0, 0, 0}}},
	}
	for _, c := range cases {
		got := playGame(c.stones, c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest playGameWithIllegalMove(stones, emptyBoard) function:")
			//fmt.Printf("stones: %v\nresult: %v\n", c.stones, c.want)
		}
	}
}

func TestCountingGroupLiberties(t *testing.T) {
	cases := []struct {
		group Stones
		board [][]int
		want int
	}{
		{Stones{{1, 1, 1}}, [][]int{[]int{0, 2, 0}, []int{2, 1, 0}, []int{0, 2, 0}}, 1},
		{Stones{{1, 1, 1}, {1, 1, 0}}, [][]int{[]int{0, 1, 0}, []int{2, 1, 0}, []int{0, 2, 0}}, 3},
		{Stones{{1, 1, 1}, {1, 1, 0}}, [][]int{[]int{0, 1, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, 5},
		{Stones{{1, 0, 0}, {1, 0, 1}, {1, 0, 2}, {1, 1, 0}, {1, 1, 2}, {1, 2, 0}, {1, 2, 1}, {1, 2, 2}}, [][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}, 1},
	}
	for _, c := range cases {
		got := countGroupLiberties(c.group, c.board)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest countGroupLiberties(group, board):")
			//fmt.Printf("result: %v\n", c.want)
		}
	}
}

func TestGroupIsAlive(t *testing.T) {
	cases := []struct {
		group Stones
		board [][]int
		want bool
	}{
		{Stones{{1, 1, 1}}, [][]int{[]int{0, 2, 0}, []int{2, 1, 2}, []int{0, 2, 0}}, false},
		{Stones{{1, 1, 1}}, [][]int{[]int{0, 0, 0}, []int{2, 1, 2}, []int{0, 2, 0}}, true},
		{Stones{{1, 1, 1}, {1, 0, 1}}, [][]int{[]int{2, 1, 2}, []int{2, 1, 2}, []int{0, 0, 0}}, true},
		{Stones{{1, 1, 1}, {1, 0, 1}}, [][]int{[]int{2, 1, 2}, []int{2, 1, 2}, []int{0, 2, 0}}, false},
		{Stones{{1, 0, 0}, {1, 0, 1}, {1, 0, 2}, {1, 1, 0}, {1, 1, 2}, {1, 2, 0}, {1, 2, 1}, {1, 2, 2}}, [][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}, true},
		{Stones{{1, 0, 0}, {1, 0, 1}, {1, 0, 2}, {1, 1, 0}, {1, 1, 2}, {1, 2, 0}, {1, 2, 1}, {1, 2, 2}}, [][]int{[]int{1, 1, 1}, []int{1, 2, 1}, []int{1, 1, 1}}, false},
	}
	for _, c := range cases {
		got := isGroupAlive(c.group, c.board)
		if !reflect.DeepEqual(got, c.want) {
			fmt.Printf("\n%v => %v\n", countGroupLiberties(c.group, c.board), findLibertiesForGroup(c.group, c.board))
			printBoardToConsole(c.board)

			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest isGroupAlive(group, board):")
			//fmt.Printf("result: %v\n", c.want)
		}
	}
}

func TestGettingMovesPartOfGameRecordAsArray(t *testing.T) {
	cases := []struct {
		gameRecord string
		want []string
	}{
		{"(;SZ[3]PB[Lee]PW[Alpha]RE[B+1];B[bb];W[ab];B[aa];W[ba])", []string{"B[bb]", "W[ab]", "B[aa]", "W[ba]"}},
	}
	for _, c := range cases {
		got := getMovesPartOfGameRecordAsArray(c.gameRecord)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Printf("Game record: %v\nMoves: %v\n", c.gameRecord, c.want)
		}
	}
}

func TestTransformingMovesToStones(t *testing.T) {
	cases := []struct {
		moves []string
		stones Stones
	}{
		{[]string{"B[bb]", "W[ab]", "B[aa]", "W[ba]"}, Stones{{1, 1, 1}, {2, 0, 1}, {1, 0, 0}, {2, 1, 0}}},
	}
	for _, c := range cases {
		got := transformMovesToStones(c.moves)
		if !reflect.DeepEqual(got, c.stones) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.stones, got)
		}
	}
}

func TestGettingMovesFromGameRecord(t *testing.T) {
	cases := []struct {
		gameRecord string
		want Stones
	}{
		{"(;SZ[3]PB[Lee]PW[Alpha]RE[B+1];B[bb];W[ab];B[aa];W[ba])", Stones{{1, 1, 1}, {2, 0, 1}, {1, 0, 0}, {2, 1, 0}}},
	}
	for _, c := range cases {
		got := getMovesFromGameRecord(c.gameRecord)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		}
	}
}

func TestWritingGameOnBoard(t *testing.T) {
	cases := []struct {
		gameRecord string
		stonesOnBoard [][]int
	}{
		{"(;SZ[3]PB[Lee]PW[Alpha]RE[B+1];B[bb];W[ab];B[aa];W[ba])", [][]int{{1, 2, 0}, {2, 1, 0}, {0, 0, 0}}},
	}
	for _, c := range cases {
		got := writeGameOnBoard(c.gameRecord)
		if !reflect.DeepEqual(got, c.stonesOnBoard) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.stonesOnBoard, got)
			printBoardToConsole(got)
		} else {
			printBoardToConsole(got)
		}
	}
}
